(define
	(fbnq n) # fbnq, a function which gain very quickly
	(cond
		((<= n 1)
			1
		)
		(1
			(loop
				(each
					(define i 0)
					(define a 1)
					(define b 1)
				)
				(< i n)
				(each
					(define i (+ i 1))
					(define c b)
					(define b (+ a b))
					(define a c)
				)
			)
		)
	)
)
(quote "ok")