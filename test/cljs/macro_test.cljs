(ns cljs.macro-test
  (:refer-clojure :exclude [==])
  (:use-macros [cljs.macro-test.macros :only [==]]))

(defn test-macros []
  (assert (= (== 1 1) 2)))

^:top-level (js*
"func Test_runner(t *testing.T) {
    ~{}
    assert.True(t, true)
}" (test-macros))
