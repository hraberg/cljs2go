(ns cljs.top-level-test)

(let [foo 1]
  (defn bar []
    foo))

(let [foo 2]
  (defn baz []
    foo))

(defn test []
  (assert (= (bar) 1))
  (assert (= (baz) 2)))

^:top-level (js*
"func Test_runner(t *testing.T) {
    ~{}
    assert.True(t, true)
}" (test))
