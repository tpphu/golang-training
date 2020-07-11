
## Load Test

## Result 1 - Before optimization

```
execution: local
    output: -
    script: k6-demo1-script.js

duration: 5s, iterations: -
        vus: 10, max: 10

done [==========================================================] 5s / 5s

✓ status was 200
✓ total should 560k

checks.....................: 100.00% ✓ 164  ✗ 0
data_received..............: 12 kB   2.5 kB/s
data_sent..................: 7.0 kB  1.4 kB/s
http_req_blocked...........: avg=295.89µs min=1µs      med=2µs      max=2.97ms p(90)=2.31ms  p(95)=2.36ms
http_req_connecting........: avg=110.87µs min=0s       med=0s       max=955µs  p(90)=868.7µs p(95)=920.9µs
http_req_duration..........: avg=554.77ms min=383.87ms med=456.37ms max=1.29s  p(90)=1.17s   p(95)=1.21s
http_req_receiving.........: avg=56.9µs   min=27µs     med=53µs     max=108µs  p(90)=73.9µs  p(95)=93.6µs
http_req_sending...........: avg=91.36µs  min=7µs      med=12µs     max=1.12ms p(90)=570.9µs p(95)=602.95µs
http_req_tls_handshaking...: avg=0s       min=0s       med=0s       max=0s     p(90)=0s      p(95)=0s
http_req_waiting...........: avg=554.62ms min=383.79ms med=456.32ms max=1.29s  p(90)=1.17s   p(95)=1.21s
http_reqs..................: 82      16.39927/s
iteration_duration.........: avg=555.12ms min=383.99ms med=456.49ms max=1.29s  p(90)=1.17s   p(95)=1.22s
iterations.................: 82      16.39927/s
vus........................: 10      min=10 max=10
vus_max....................: 10      min=10 max=10
```

## Result 2 - After optimization

Expection: `iterations.................: 160      30.39927/s` (x2)
Actually:  `iterations.................: 233     46.598695/s` (x3)

```
execution: local
    output: -
    script: k6-demo1-script.js

duration: 5s, iterations: -
        vus: 10, max: 10

done [==========================================================] 5s / 5s

✓ status was 200
✓ total should 560k

checks.....................: 100.00% ✓ 466  ✗ 0
data_received..............: 35 kB   7.1 kB/s
data_sent..................: 20 kB   4.0 kB/s
http_req_blocked...........: avg=60.51µs  min=1µs      med=2µs      max=1.51ms   p(90)=3.8µs    p(95)=10.39µs
http_req_connecting........: avg=21µs     min=0s       med=0s       max=581µs    p(90)=0s       p(95)=0s
http_req_duration..........: avg=210.4ms  min=129.83ms med=189.49ms max=637.85ms p(90)=250.73ms p(95)=291.29ms
http_req_receiving.........: avg=56.15µs  min=25µs     med=52µs     max=219µs    p(90)=72µs     p(95)=85.19µs
http_req_sending...........: avg=17.7µs   min=6µs      med=12µs     max=169µs    p(90)=22µs     p(95)=32.59µs
http_req_tls_handshaking...: avg=0s       min=0s       med=0s       max=0s       p(90)=0s       p(95)=0s
http_req_waiting...........: avg=210.33ms min=129.77ms med=189.44ms max=637.67ms p(90)=250.65ms p(95)=291.23ms
http_reqs..................: 233     46.598695/s
iteration_duration.........: avg=210.59ms min=129.96ms med=189.57ms max=639.27ms p(90)=250.9ms  p(95)=291.39ms
iterations.................: 233     46.598695/s
vus........................: 10      min=10 max=10
vus_max....................: 10      min=10 max=10
```

## Issues

```go
func total() int {
	urls := []string{
		"https://www.sendo.vn/m/wap_v2/full/san-pham/ao-so-mi-jean-nam-dai-tay-cao-cap-hang-vnxk-31331127?platform=web",
		"https://www.sendo.vn/m/wap_v2/full/san-pham/ao-dui-nam-cao-cap-30157047",
		"https://www.sendo.vn/m/wap_v2/full/san-pham/ao-so-mi-nam-hang-hop-10036141"}

	total := 0
	wg := sync.WaitGroup{} // wait
	for _, url := range urls {
		wg.Add(1) // wait
		// PV nhieu cong ty dung: Nodejs/Go/Java
		go func() {
			defer wg.Done() //wait
			product, err := getProduct(url)
			if err != nil {
				panic(err)
			}
			total = total + product.Price
		}()
	}
	wg.Wait() //wait
	return total
}
```

```
execution: local
output: -
script: k6-demo1-script.js

duration: 5s, iterations: -
    vus: 10, max: 10

done [==========================================================] 5s / 5s

✓ status was 200
✗ total should 560k
↳  0% — ✓ 0 / ✗ 51

checks.....................: 50.00% ✓ 51   ✗ 51
data_received..............: 7.8 kB 1.6 kB/s
data_sent..................: 4.3 kB 866 B/s
http_req_blocked...........: avg=337.25µs min=1µs      med=2µs      max=2.23ms p(90)=1.74ms p(95)=1.96ms
http_req_connecting........: avg=125.98µs min=0s       med=0s       max=903µs  p(90)=714µs  p(95)=771.5µs
http_req_duration..........: avg=906.85ms min=213.15ms med=538.63ms max=2.44s  p(90)=2.44s  p(95)=2.44s
http_req_receiving.........: avg=91.56µs  min=27µs     med=61µs     max=850µs  p(90)=104µs  p(95)=246.5µs
http_req_sending...........: avg=90.35µs  min=7µs      med=15µs     max=618µs  p(90)=359µs  p(95)=526.99µs
http_req_tls_handshaking...: avg=0s       min=0s       med=0s       max=0s     p(90)=0s     p(95)=0s
http_req_waiting...........: avg=906.67ms min=213.1ms  med=538.25ms max=2.44s  p(90)=2.44s  p(95)=2.44s
http_reqs..................: 51     10.199878/s
iteration_duration.........: avg=907.33ms min=213.25ms med=538.76ms max=2.44s  p(90)=2.44s  p(95)=2.44s
iterations.................: 51     10.199878/s
vus........................: 10     min=10 max=10
vus_max....................: 10     min=10 max=10
```