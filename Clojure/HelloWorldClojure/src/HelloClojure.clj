(ns HelloClojure)

(defn fib
  "docstring"
  ([n]
   (fib [0 1] n))
  ([x, n]
   (if (< (count x) n)
   (fib (conj x (+ (last x) (nth x (- (count x) 2)))) n) x)))

(defn add
  [x y]
  (+ x y))



(print "Hello World, clojure\n")
(println(fib 10))
(println (add 1 12))