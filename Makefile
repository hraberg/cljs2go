# LAKE - Leiningen AKcelErator

CLASSPATH=$(shell cat .lein-classpath)
LEIN_MAIN=$(shell cat project.clj | grep :main | awk '{ print $$2 }')
LEIN_JVM_OPTS=-Xms512m -Xmx512m -noverify -XX:+TieredCompilation -XX:TieredStopAtLevel=1
CLJ=java $(JVM_OPTS) $(LEIN_JVM_OPTS) -Xbootclasspath/a:$(CLASSPATH) clojure.main
LEIN_REPL_PORT=0
LEIN_REPL_HOST="127.0.0.1"
LEIN_REPL_TRANSPORT=clojure.tools.nrepl.transport/bencode
LEIN_REPL_HANDLER=clojure.tools.nrepl.server/default-handler
WATCH=$(shell echo $(CLASSPATH) | tr : ' ' | sed -E 's/\S+\.jar ?//g' | xargs ls -d {} 2> /dev/null)
LEIN=$(shell (which lein || echo ./lein))

lein:
	wget -O lein https://raw.githubusercontent.com/technomancy/leiningen/stable/bin/lein
	chmod +x lein

.lein-classpath: project.clj $(LEIN)
	$(LEIN) classpath > .lein-classpath

classpath: .lein-classpath

define RLWRAP_COMPLETIONS
(require (symbol "clojure.tools.namespace.find"))
(doseq [n (->> (.split (str (System/getProperty "sun.boot.class.path") ":"
                            (System/getProperty "java.class.path"))
                        ":")
               (map clojure.java.io/file)
               clojure.tools.namespace.find/find-namespaces)]
  (try (require n) (catch Exception _)))
(->> (for [n (all-ns) [k v] (ns-publics n)]
        [(ns-name n) k (subs (str v) 2)])
	 (reduce into (hash-set))
     (apply print-str)
     (spit ".rlwrap-completions"))
(System/exit 0)
endef
export RLWRAP_COMPLETIONS

.rlwrap-completions: .lein-classpath
	@$(CLJ) -e "$$RLWRAP_COMPLETIONS"

rlwrap-repl: classpath .rlwrap-completions
	@rlwrap --prompt-colour=Red --multi-line --remember --complete-filenames -b "(){}[],^%$#@\"\";:''|\\" \
	  -f .rlwrap-completions $(CLJ)

define REPL
(require (symbol "clojure.tools.nrepl.server"))
(-> "$(LEIN_REPL_HANDLER)" symbol namespace symbol require)
(let [{:keys [ss port]} (clojure.tools.nrepl.server/start-server :host $(LEIN_REPL_HOST)
                                                                 :port $(LEIN_REPL_PORT)
                                                                 :transport-fn $(LEIN_REPL_TRANSPORT)
                                                                 :handler (if (= clojure.tools.nrepl.server/default-handler
                                                                                 $(LEIN_REPL_HANDLER))
                                                                             ($(LEIN_REPL_HANDLER))
                                                                             $(LEIN_REPL_HANDLER)))
      host (.getHostName (.getInetAddress ^java.net.ServerSocket ss))]
      (printf  "nREPL server started on port %s on host %s - nrepl://%s:%s\n" port host host port))
endef
export REPL

# cider-jack-in
repl: classpath
	@$(CLJ) -e "$$REPL"

define TEST
(require (symbol "clojure.test"))
(require (symbol "clojure.tools.namespace.repl"))
(clojure.tools.namespace.repl/refresh)
(clojure.test/run-all-tests #"^(?!clojure|user).*test.*")
endef
export TEST

run: classpath
	@$(CLJ) -m $(LEIN_MAIN)

test: classpath
	@$(CLJ) -e "$$TEST"

check: test

clean:
	rm -rf target

inotifywait:
ifeq ($(shell which inotifywait),)
	sudo apt-get install inotify-tools
endif

test-refresh: classpath inotifywait
	@bash -c '(cat <&0 & inotifywait -qmre create -e modify --exclude .+~ --exclude \.# $(WATCH)) \
	            | $(CLJ) -e "(do $$TEST (while (read-line) $$TEST))"'

bench-clojure.main: classpath
	bash -c 'time echo "(System/exit 0)" | $(CLJ)'

bench-repl: classpath
	bash -c 'time $(CLJ) -e "(do $$REPL (flush) (System/exit 0))"'
