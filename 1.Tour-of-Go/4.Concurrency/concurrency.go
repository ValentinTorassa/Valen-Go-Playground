// go_concurrency_notes.go
// Resumen práctico (con código y comentarios) de Goroutines, Channels, Select, Mutex y Crawler concurrente.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("== 1) Goroutines básicas ==")
	goroutinesBasics()

	fmt.Println("\n== 2) Channels sin buffer (sincronización implícita) ==")
	channelsUnbuffered()

	fmt.Println("\n== 3) Channels con buffer (capacidad y bloqueo) ==")
	channelsBuffered()

	fmt.Println("\n== 4) Range y close sobre channels ==")
	rangeAndClose()

	fmt.Println("\n== 5) Select con default (no-bloqueante) y multiplexación ==")
	selectExamples()

	fmt.Println("\n== 6) sync.Mutex (exclusión mutua sin comunicación) ==")
	mutexExample()

	fmt.Println("\n== 7) Ejercicio: Web Crawler concurrente con cache seguro ==")
	crawlerDemo()
}

// ------------------------------------------------------------
// 1) GOROUTINES
// Una goroutine es un hilo ligero gestionado por el runtime de Go.
// "go f(x,y,z)" evalúa f,x,y,z en la goroutine actual y ejecuta f(...) en una goroutine nueva.
// Comparten espacio de direcciones → ¡sincronizar accesos compartidos!
// ------------------------------------------------------------
func goroutinesBasics() {
	work := func(id int) {
		fmt.Printf("[worker %d] start\n", id)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		fmt.Printf("[worker %d] done\n", id)
	}
	for i := 1; i <= 3; i++ {
		go work(i) // lanzar 3 goroutines
	}
	// Espera “best effort” para ver la salida (en código real usarías sync.WaitGroup)
	time.Sleep(400 * time.Millisecond)
}

// ------------------------------------------------------------
// 2) CHANNELS (tipados): envían/reciben valores con "<-"
// Por defecto, send/receive BLOQUEAN hasta que el otro lado esté listo.
// ch := make(chan int) crea un channel sin buffer.
// ------------------------------------------------------------
func channelsUnbuffered() {
	sum := func(nums []int, out chan<- int) { // "chan<- int" = solo enviar
		total := 0
		for _, n := range nums {
			total += n
		}
		out <- total // bloquea hasta que alguien reciba
	}

	ch := make(chan int) // sin buffer
	go sum([]int{1, 2, 3, 4, 5}, ch)
	go sum([]int{6, 7, 8, 9}, ch)

	// Recibir de a uno: cada Receive bloquea hasta que haya envío
	a := <-ch
	b := <-ch
	fmt.Println("partial sums:", a, b, "→ final:", a+b)
}

// ------------------------------------------------------------
// 3) CHANNELS CON BUFFER: ch := make(chan T, N)
// Envíos bloquean solo si el buffer está LLENO; recepciones si está VACÍO.
// ------------------------------------------------------------
func channelsBuffered() {
	ch := make(chan string, 2) // buffer de 2
	ch <- "A"
	ch <- "B"
	// ch <- "C" // ← si descomentas, BLOQUEA aquí hasta que alguien reciba (buffer lleno)

	fmt.Println(<-ch) // "A"
	fmt.Println(<-ch) // "B"
	// Ya hay espacio; ahora enviar no bloquea:
	ch <- "C"
	fmt.Println(<-ch) // "C"
}

// ------------------------------------------------------------
// 4) RANGE y CLOSE:
// - Solo el EMISOR debe close(ch).
// - for v := range ch lee hasta que el channel se cierra.
// - Recepción con "ok" indica si el channel está cerrado (ok=false).
// ------------------------------------------------------------
func rangeAndClose() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) // señal: no habrá más valores
	}()
	for v := range ch { // itera hasta close
		fmt.Print(v, " ")
	}
	fmt.Println()

	// Recepción con "ok":
	ch2 := make(chan int, 1)
	ch2 <- 42
	close(ch2)
	if v, ok := <-ch2; ok {
		fmt.Println("got:", v)
	}
	if _, ok := <-ch2; !ok {
		fmt.Println("channel closed (no more values)")
	}
}

// ------------------------------------------------------------
// 5) SELECT: espera múltiples operaciones de comunicación.
// - Se bloquea hasta que algún case esté listo; si varios, elige uno al azar.
// - "default" corre si ningún case está listo (try-send/try-recv no bloqueante).
// ------------------------------------------------------------
func selectExamples() {
	// Multiplexar dos productores
	a := make(chan string)
	b := make(chan string)
	go func() {
		time.Sleep(60 * time.Millisecond)
		a <- "from A"
	}()
	go func() {
		time.Sleep(30 * time.Millisecond)
		b <- "from B"
	}()

	select {
	case v := <-a:
		fmt.Println("select got:", v)
	case v := <-b:
		fmt.Println("select got:", v)
	}

	// Intento no bloqueante con default
	c := make(chan int, 1)
	select {
	case c <- 99:
		fmt.Println("sent 99 without blocking")
	default:
		fmt.Println("send would block")
	}

	select {
	case v := <-c:
		fmt.Println("recv:", v)
	default:
		fmt.Println("recv would block")
	}
}

// ------------------------------------------------------------
// 6) MUTEX (exclusión mutua)
// Útil cuando NO se necesita comunicación, solo proteger datos compartidos.
// sync.Mutex tiene Lock/Unlock; usar defer Unlock para no olvidarlo.
// ------------------------------------------------------------
func mutexExample() {
	var (
		mu   sync.Mutex
		cnt  int
		wg   sync.WaitGroup
		nG   = 4
		loop = 1_000
	)

	inc := func() {
		defer wg.Done()
		for i := 0; i < loop; i++ {
			mu.Lock()
			cnt++
			mu.Unlock()
		}
	}

	wg.Add(nG)
	for i := 0; i < nG; i++ {
		go inc()
	}
	wg.Wait()
	fmt.Println("counter =", cnt, "(esperado:", nG*loop, ")")
}

// ------------------------------------------------------------
// 7) Web Crawler concurrente
// Objetivo: crawlear URLs en paralelo SIN repetir (cache concurrente).
// - Usamos un "fetch" simulado (map) para no depender de red.
// - visitedMap + Mutex para seguridad concurrente.
// - WaitGroup para esperar a que terminen las goroutines.
// NOTA: maps "solos" NO son seguros para concurrencia; por eso usamos Mutex.
// ------------------------------------------------------------
func crawlerDemo() {
	// Fake web: URL -> links
	web := map[string][]string{
		"https://golang.org/": {
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
		"https://golang.org/pkg/": {
			"https://golang.org/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
		"https://golang.org/pkg/fmt/": {
			"https://golang.org/",
		},
		"https://golang.org/pkg/os/": {
			"https://golang.org/",
		},
	}

	fetch := func(url string) ([]string, error) {
		time.Sleep(20 * time.Millisecond) // simula latencia
		if links, ok := web[url]; ok {
			return links, nil
		}
		return nil, fmt.Errorf("not found: %s", url)
	}

	type SafeSet struct {
		mu sync.Mutex
		m  map[string]struct{}
	}

	seen := &SafeSet{m: make(map[string]struct{})}
	markSeen := func(url string) bool {
		// true si marcamos por primera vez; false si ya existía
		seen.mu.Lock()
		defer seen.mu.Unlock()
		if _, ok := seen.m[url]; ok {
			return false
		}
		seen.m[url] = struct{}{}
		return true
	}

	var wg sync.WaitGroup

	var crawl func(string, int)
	crawl = func(url string, depth int) {
		defer wg.Done()
		if depth <= 0 {
			return
		}
		if !markSeen(url) {
			return // ya se visitó
		}
		fmt.Println("fetch:", url)

		links, err := fetch(url)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		// Lanzar nuevas goroutines para los links
		for _, u := range links {
			wg.Add(1)
			go crawl(u, depth-1)
		}
	}

	start := "https://golang.org/"
	wg.Add(1)
	go crawl(start, 3)
	wg.Wait()

	fmt.Println("visited:")
	for url := range seen.m {
		fmt.Println(" -", url)
	}
}

/*
RESUMEN CONCEPTUAL (para pegar en tu editor):
- Goroutines: go f(...) crea un hilo ligero. Comparten memoria → sincronizar.
- Channels: make(chan T[, n]); envíos/recepciones bloquean por defecto. Con buffer n: send bloquea si lleno; recv si vacío.
- Range/Close: solo la parte emisora debe close(ch). range lee hasta que se cierre.
- Select: espera múltiples operaciones de channel; default evita bloqueo.
- Mutex: usa Lock/Unlock para proteger estado compartido cuando no hace falta comunicar.
- Crawler: paraleliza el fetch de URLs y evita duplicados con un map protegido por Mutex.
*/
