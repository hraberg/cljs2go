// Compiled by ClojureScript to Go 0.0-2322
// clojure.zip

// Functional hierarchical zipper, with navigation, editing,
// and enumeration.  See Huet
// Author: Rich Hickey
package zip

import cljs_core "github.com/hraberg/cljs.go/cljs/core"

// Creates a new zipper structure.
//
// branch? is a fn that, given a node, returns true if can have
// children, even if it currently doesn't.
//
// children is a fn that, given a branch node, returns a seq of its
// children.
//
// make-node is a fn that, given an existing node and a seq of
// children, returns a new branch node with the supplied children.
// root is the root node.
var Zipper *cljs_core.AFn

func init() {
	Zipper = func(zipper *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(zipper, func(branch_QMARK_ interface{}, children interface{}, make_node interface{}, root interface{}) interface{} {
			return cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{root, nil}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, 3, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "zip", Name: "make-node", Fqn: "zip/make-node", X_hash: float64(1103800591)}), make_node, (&cljs_core.CljsCoreKeyword{Ns: "zip", Name: "children", Fqn: "zip/children", X_hash: float64(-940194589)}), children, (&cljs_core.CljsCoreKeyword{Ns: "zip", Name: "branch?", Fqn: "zip/branch?", X_hash: float64(-998880862)}), branch_QMARK_}, nil}))
		})
	}(&cljs_core.AFn{})
}

// Returns a zipper for nested sequences, given a root sequence
var Seq_zip *cljs_core.AFn

func init() {
	Seq_zip = func(seq_zip *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(seq_zip, func(root interface{}) interface{} {
			return Zipper.X_invoke_Arity4(cljs_core.Seq_QMARK_, cljs_core.Identity, func(G__1 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1, func(node interface{}, children interface{}) interface{} {
					return cljs_core.With_meta.X_invoke_Arity2(children, cljs_core.Meta.X_invoke_Arity1(node))
				})
			}(&cljs_core.AFn{}), root)
		})
	}(&cljs_core.AFn{})
}

// Returns a zipper for nested vectors, given a root vector
var Vector_zip *cljs_core.AFn

func init() {
	Vector_zip = func(vector_zip *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(vector_zip, func(root interface{}) interface{} {
			return Zipper.X_invoke_Arity4(cljs_core.Vector_QMARK_, cljs_core.Seq, func(G__2 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__2, func(node interface{}, children interface{}) interface{} {
					return cljs_core.With_meta.X_invoke_Arity2(cljs_core.Vec.X_invoke_Arity1(children), cljs_core.Meta.X_invoke_Arity1(node))
				})
			}(&cljs_core.AFn{}), root)
		})
	}(&cljs_core.AFn{})
}

// Returns a zipper for xml elements (as from xml/parse),
// given a root element
var Xml_zip *cljs_core.AFn

func init() {
	Xml_zip = func(xml_zip *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(xml_zip, func(root interface{}) interface{} {
			return Zipper.X_invoke_Arity4(cljs_core.Complement.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Comp.X_invoke_Arity2(cljs_core.Seq, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "content", Fqn: "content", X_hash: float64(15833224)})).(cljs_core.CljsCoreIFn), func(G__3 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__3, func(node interface{}, children interface{}) interface{} {
					return cljs_core.Assoc.X_invoke_Arity3(node, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "content", Fqn: "content", X_hash: float64(15833224)}), func() interface{} {
						var and__79525__auto__ = children
						_ = and__79525__auto__
						if cljs_core.Truth_(and__79525__auto__) {
							return cljs_core.Apply.X_invoke_Arity2(cljs_core.Vector, children)
						} else {
							return and__79525__auto__
						}
					}())
				})
			}(&cljs_core.AFn{}), root)
		})
	}(&cljs_core.AFn{})
}

// Returns the node at loc
var Node *cljs_core.AFn

func init() {
	Node = func(node *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(node, func(loc interface{}) interface{} {
			return loc.(cljs_core.CljsCoreIFn).X_invoke_Arity1(float64(0))
		})
	}(&cljs_core.AFn{})
}

// Returns true if the node at loc is a branch
var Branch_QMARK_ *cljs_core.AFn

func init() {
	Branch_QMARK_ = func(branch_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(branch_QMARK_, func(loc interface{}) interface{} {
			return (&cljs_core.CljsCoreKeyword{Ns: "zip", Name: "branch?", Fqn: "zip/branch?", X_hash: float64(-998880862)}).X_invoke_Arity1(cljs_core.Meta.X_invoke_Arity1(loc)).(cljs_core.CljsCoreIFn).X_invoke_Arity1(Node.X_invoke_Arity1(loc))
		})
	}(&cljs_core.AFn{})
}

// Returns a seq of the children of node at loc, which must be a branch
var Children *cljs_core.AFn

func init() {
	Children = func(children *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(children, func(loc interface{}) interface{} {
			if cljs_core.Truth_(Branch_QMARK_.X_invoke_Arity1(loc)) {
				return (&cljs_core.CljsCoreKeyword{Ns: "zip", Name: "children", Fqn: "zip/children", X_hash: float64(-940194589)}).X_invoke_Arity1(cljs_core.Meta.X_invoke_Arity1(loc)).(cljs_core.CljsCoreIFn).X_invoke_Arity1(Node.X_invoke_Arity1(loc))
			} else {
				panic("called children on a leaf node")
			}
		})
	}(&cljs_core.AFn{})
}

// Returns a new branch node, given an existing node and new
// children. The loc is only used to supply the constructor.
var Make_node *cljs_core.AFn

func init() {
	Make_node = func(make_node *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(make_node, func(loc interface{}, node interface{}, children interface{}) interface{} {
			return (&cljs_core.CljsCoreKeyword{Ns: "zip", Name: "make-node", Fqn: "zip/make-node", X_hash: float64(1103800591)}).X_invoke_Arity1(cljs_core.Meta.X_invoke_Arity1(loc)).(cljs_core.CljsCoreIFn).X_invoke_Arity2(node, children)
		})
	}(&cljs_core.AFn{})
}

// Returns a seq of nodes leading to this loc
var Path *cljs_core.AFn

func init() {
	Path = func(path *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(path, func(loc interface{}) interface{} {
			return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "pnodes", Fqn: "pnodes", X_hash: float64(1739080565)}).X_invoke_Arity1(loc.(cljs_core.CljsCoreIFn).X_invoke_Arity1(float64(1)))
		})
	}(&cljs_core.AFn{})
}

// Returns a seq of the left siblings of this loc
var Lefts *cljs_core.AFn

func init() {
	Lefts = func(lefts *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(lefts, func(loc interface{}) interface{} {
			return cljs_core.Seq.Arity1IQ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}).X_invoke_Arity1(loc.(cljs_core.CljsCoreIFn).X_invoke_Arity1(float64(1))))
		})
	}(&cljs_core.AFn{})
}

// Returns a seq of the right siblings of this loc
var Rights *cljs_core.AFn

func init() {
	Rights = func(rights *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(rights, func(loc interface{}) interface{} {
			return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}).X_invoke_Arity1(loc.(cljs_core.CljsCoreIFn).X_invoke_Arity1(float64(1)))
		})
	}(&cljs_core.AFn{})
}

// Returns the loc of the leftmost child of the node at this loc, or
// nil if no children
var Down *cljs_core.AFn

func init() {
	Down = func(down *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(down, func(loc interface{}) interface{} {
			if cljs_core.Truth_(Branch_QMARK_.X_invoke_Arity1(loc)) {
				{
					var vec__6 = loc
					var node = cljs_core.Nth.X_invoke_Arity3(vec__6, float64(0), nil)
					var path = cljs_core.Nth.X_invoke_Arity3(vec__6, float64(1), nil)
					var vec__7 = Children.X_invoke_Arity1(loc)
					var c = cljs_core.Nth.X_invoke_Arity3(vec__7, float64(0), nil)
					var cnext = cljs_core.Seq_(cljs_core.Nthnext.X_invoke_Arity2(vec__7, float64(1)))
					var cs = vec__7
					_, _, _, _, _, _, _ = vec__6, node, path, vec__7, c, cnext, cs
					if cljs_core.Truth_(cs) {
						return cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{c, (&cljs_core.CljsCorePersistentArrayMap{nil, 4, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}), cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "pnodes", Fqn: "pnodes", X_hash: float64(1739080565)}), func() interface{} {
							if cljs_core.Truth_(path) {
								return cljs_core.Conj.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "pnodes", Fqn: "pnodes", X_hash: float64(1739080565)}).X_invoke_Arity1(path), node)
							} else {
								return (&cljs_core.CljsCorePersistentVector{nil, 1, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{node}, nil})
							}
						}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ppath", Fqn: "ppath", X_hash: float64(-1758182784)}), path, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}), cnext}, nil})}, nil}), cljs_core.Meta.X_invoke_Arity1(loc))
					} else {
						return nil
					}
				}
			} else {
				return nil
			}
		})
	}(&cljs_core.AFn{})
}

// Returns the loc of the parent of the node at this loc, or nil if at
// the top
var Up *cljs_core.AFn

func init() {
	Up = func(up *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(up, func(loc interface{}) interface{} {
			{
				var vec__10 = loc
				var node = cljs_core.Nth.X_invoke_Arity3(vec__10, float64(0), nil)
				var map__11 = cljs_core.Nth.X_invoke_Arity3(vec__10, float64(1), nil)
				var map__11___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__11) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__11)
					} else {
						return map__11
					}
				}()
				var path = map__11___1
				var l = cljs_core.Get.X_invoke_Arity2(map__11___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}))
				var ppath = cljs_core.Get.X_invoke_Arity2(map__11___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ppath", Fqn: "ppath", X_hash: float64(-1758182784)}))
				var pnodes = cljs_core.Get.X_invoke_Arity2(map__11___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "pnodes", Fqn: "pnodes", X_hash: float64(1739080565)}))
				var r = cljs_core.Get.X_invoke_Arity2(map__11___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}))
				var changed_QMARK_ = cljs_core.Get.X_invoke_Arity2(map__11___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "changed?", Fqn: "changed?", X_hash: float64(-437828330)}))
				_, _, _, _, _, _, _, _, _, _ = vec__10, node, map__11, map__11___1, path, l, ppath, pnodes, r, changed_QMARK_
				if cljs_core.Truth_(pnodes) {
					{
						var pnode = cljs_core.Peek.X_invoke_Arity1(pnodes)
						_ = pnode
						return cljs_core.With_meta.X_invoke_Arity2(func() cljs_core.CljsCoreIVector {
							if cljs_core.Truth_(changed_QMARK_) {
								return (&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{Make_node.X_invoke_Arity3(loc, pnode, cljs_core.Concat.X_invoke_Arity2(l, cljs_core.Cons.X_invoke_Arity2(node, r).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreLazySeq)), func() interface{} {
									var and__79525__auto__ = ppath
									_ = and__79525__auto__
									if cljs_core.Truth_(and__79525__auto__) {
										return cljs_core.Assoc.X_invoke_Arity3(ppath, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "changed?", Fqn: "changed?", X_hash: float64(-437828330)}), true)
									} else {
										return and__79525__auto__
									}
								}()}, nil})
							} else {
								return (&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{pnode, ppath}, nil})
							}
						}(), cljs_core.Meta.X_invoke_Arity1(loc))
					}
				} else {
					return nil
				}
			}
		})
	}(&cljs_core.AFn{})
}

// zips all the way up and returns the root node, reflecting any
// changes.
var Root *cljs_core.AFn

func init() {
	Root = func(root *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(root, func(loc interface{}) interface{} {
			for {
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end", Fqn: "end", X_hash: float64(-268185958)}), loc.(cljs_core.CljsCoreIFn).X_invoke_Arity1(float64(1))) {
					return Node.X_invoke_Arity1(loc)
				} else {
					{
						var p = Up.X_invoke_Arity1(loc)
						_ = p
						if cljs_core.Truth_(p) {
							loc = p
							continue
						} else {
							return Node.X_invoke_Arity1(loc)
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})
}

// Returns the loc of the right sibling of the node at this loc, or nil
var Right *cljs_core.AFn

func init() {
	Right = func(right *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(right, func(loc interface{}) interface{} {
			{
				var vec__15 = loc
				var node = cljs_core.Nth.X_invoke_Arity3(vec__15, float64(0), nil)
				var map__16 = cljs_core.Nth.X_invoke_Arity3(vec__15, float64(1), nil)
				var map__16___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__16) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__16)
					} else {
						return map__16
					}
				}()
				var path = map__16___1
				var l = cljs_core.Get.X_invoke_Arity2(map__16___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}))
				var vec__17 = cljs_core.Get.X_invoke_Arity2(map__16___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}))
				var r = cljs_core.Nth.X_invoke_Arity3(vec__17, float64(0), nil)
				var rnext = cljs_core.Seq_(cljs_core.Nthnext.X_invoke_Arity2(vec__17, float64(1)))
				var rs = vec__17
				_, _, _, _, _, _, _, _, _, _ = vec__15, node, map__16, map__16___1, path, l, vec__17, r, rnext, rs
				if cljs_core.Truth_(func() interface{} {
					var and__79525__auto__ = path
					_ = and__79525__auto__
					if cljs_core.Truth_(and__79525__auto__) {
						return rs
					} else {
						return and__79525__auto__
					}
				}()) {
					return cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{r, cljs_core.Assoc.X_invoke_ArityVariadic(path, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}), cljs_core.Conj.X_invoke_Arity2(l, node), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}), rnext)}, nil}), cljs_core.Meta.X_invoke_Arity1(loc))
				} else {
					return nil
				}
			}
		})
	}(&cljs_core.AFn{})
}

// Returns the loc of the rightmost sibling of the node at this loc, or self
var Rightmost *cljs_core.AFn

func init() {
	Rightmost = func(rightmost *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(rightmost, func(loc interface{}) interface{} {
			{
				var vec__20 = loc
				var node = cljs_core.Nth.X_invoke_Arity3(vec__20, float64(0), nil)
				var map__21 = cljs_core.Nth.X_invoke_Arity3(vec__20, float64(1), nil)
				var map__21___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__21) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__21)
					} else {
						return map__21
					}
				}()
				var path = map__21___1
				var l = cljs_core.Get.X_invoke_Arity2(map__21___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}))
				var r = cljs_core.Get.X_invoke_Arity2(map__21___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}))
				_, _, _, _, _, _, _ = vec__20, node, map__21, map__21___1, path, l, r
				if cljs_core.Truth_(func() interface{} {
					var and__79525__auto__ = path
					_ = and__79525__auto__
					if cljs_core.Truth_(and__79525__auto__) {
						return r
					} else {
						return and__79525__auto__
					}
				}()) {
					return cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.Last.X_invoke_Arity1(r), cljs_core.Assoc.X_invoke_ArityVariadic(path, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}), cljs_core.Apply.X_invoke_Arity4(cljs_core.Conj, l, node, cljs_core.Seq_(cljs_core.Butlast.X_invoke_Arity1(r))), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}), nil)}, nil}), cljs_core.Meta.X_invoke_Arity1(loc))
				} else {
					return loc
				}
			}
		})
	}(&cljs_core.AFn{})
}

// Returns the loc of the left sibling of the node at this loc, or nil
var Left *cljs_core.AFn

func init() {
	Left = func(left *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(left, func(loc interface{}) interface{} {
			{
				var vec__24 = loc
				var node = cljs_core.Nth.X_invoke_Arity3(vec__24, float64(0), nil)
				var map__25 = cljs_core.Nth.X_invoke_Arity3(vec__24, float64(1), nil)
				var map__25___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__25) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__25)
					} else {
						return map__25
					}
				}()
				var path = map__25___1
				var l = cljs_core.Get.X_invoke_Arity2(map__25___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}))
				var r = cljs_core.Get.X_invoke_Arity2(map__25___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}))
				_, _, _, _, _, _, _ = vec__24, node, map__25, map__25___1, path, l, r
				if cljs_core.Truth_(func() interface{} {
					var and__79525__auto__ = path
					_ = and__79525__auto__
					if cljs_core.Truth_(and__79525__auto__) {
						return cljs_core.Seq.Arity1IQ(l)
					} else {
						return and__79525__auto__
					}
				}()) {
					return cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.Peek.X_invoke_Arity1(l), cljs_core.Assoc.X_invoke_ArityVariadic(path, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}), cljs_core.Pop.X_invoke_Arity1(l), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}), cljs_core.Cons.X_invoke_Arity2(node, r).(*cljs_core.CljsCoreCons))}, nil}), cljs_core.Meta.X_invoke_Arity1(loc))
				} else {
					return nil
				}
			}
		})
	}(&cljs_core.AFn{})
}

// Returns the loc of the leftmost sibling of the node at this loc, or self
var Leftmost *cljs_core.AFn

func init() {
	Leftmost = func(leftmost *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(leftmost, func(loc interface{}) interface{} {
			{
				var vec__28 = loc
				var node = cljs_core.Nth.X_invoke_Arity3(vec__28, float64(0), nil)
				var map__29 = cljs_core.Nth.X_invoke_Arity3(vec__28, float64(1), nil)
				var map__29___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__29) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__29)
					} else {
						return map__29
					}
				}()
				var path = map__29___1
				var l = cljs_core.Get.X_invoke_Arity2(map__29___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}))
				var r = cljs_core.Get.X_invoke_Arity2(map__29___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}))
				_, _, _, _, _, _, _ = vec__28, node, map__29, map__29___1, path, l, r
				if cljs_core.Truth_(func() interface{} {
					var and__79525__auto__ = path
					_ = and__79525__auto__
					if cljs_core.Truth_(and__79525__auto__) {
						return cljs_core.Seq.Arity1IQ(l)
					} else {
						return and__79525__auto__
					}
				}()) {
					return cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.First.X_invoke_Arity1(l), cljs_core.Assoc.X_invoke_ArityVariadic(path, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}), cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}), cljs_core.Concat.X_invoke_ArityVariadic(cljs_core.Rest.Arity1IQ(l), (&cljs_core.CljsCorePersistentVector{nil, 1, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{node}, nil}), r).(*cljs_core.CljsCoreLazySeq))}, nil}), cljs_core.Meta.X_invoke_Arity1(loc))
				} else {
					return loc
				}
			}
		})
	}(&cljs_core.AFn{})
}

// Inserts the item as the left sibling of the node at this loc,
// without moving
var Insert_left *cljs_core.AFn

func init() {
	Insert_left = func(insert_left *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(insert_left, func(loc interface{}, item interface{}) interface{} {
			{
				var vec__32 = loc
				var node = cljs_core.Nth.X_invoke_Arity3(vec__32, float64(0), nil)
				var map__33 = cljs_core.Nth.X_invoke_Arity3(vec__32, float64(1), nil)
				var map__33___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__33) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__33)
					} else {
						return map__33
					}
				}()
				var path = map__33___1
				var l = cljs_core.Get.X_invoke_Arity2(map__33___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}))
				_, _, _, _, _, _ = vec__32, node, map__33, map__33___1, path, l
				if cljs_core.Nil_(path) {
					panic("Insert at top")
				} else {
					return cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{node, cljs_core.Assoc.X_invoke_ArityVariadic(path, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}), cljs_core.Conj.X_invoke_Arity2(l, item), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "changed?", Fqn: "changed?", X_hash: float64(-437828330)}), true)}, nil}), cljs_core.Meta.X_invoke_Arity1(loc))
				}
			}
		})
	}(&cljs_core.AFn{})
}

// Inserts the item as the right sibling of the node at this loc,
// without moving
var Insert_right *cljs_core.AFn

func init() {
	Insert_right = func(insert_right *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(insert_right, func(loc interface{}, item interface{}) interface{} {
			{
				var vec__36 = loc
				var node = cljs_core.Nth.X_invoke_Arity3(vec__36, float64(0), nil)
				var map__37 = cljs_core.Nth.X_invoke_Arity3(vec__36, float64(1), nil)
				var map__37___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__37) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__37)
					} else {
						return map__37
					}
				}()
				var path = map__37___1
				var r = cljs_core.Get.X_invoke_Arity2(map__37___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}))
				_, _, _, _, _, _ = vec__36, node, map__37, map__37___1, path, r
				if cljs_core.Nil_(path) {
					panic("Insert at top")
				} else {
					return cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{node, cljs_core.Assoc.X_invoke_ArityVariadic(path, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}), cljs_core.Cons.X_invoke_Arity2(item, r).(*cljs_core.CljsCoreCons), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "changed?", Fqn: "changed?", X_hash: float64(-437828330)}), true)}, nil}), cljs_core.Meta.X_invoke_Arity1(loc))
				}
			}
		})
	}(&cljs_core.AFn{})
}

// Replaces the node at this loc, without moving
var Replace *cljs_core.AFn

func init() {
	Replace = func(replace *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(replace, func(loc interface{}, node interface{}) interface{} {
			{
				var vec__39 = loc
				var ___ = cljs_core.Nth.X_invoke_Arity3(vec__39, float64(0), nil)
				var path = cljs_core.Nth.X_invoke_Arity3(vec__39, float64(1), nil)
				_, _, _ = vec__39, ___, path
				return cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{node, cljs_core.Assoc.X_invoke_Arity3(path, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "changed?", Fqn: "changed?", X_hash: float64(-437828330)}), true)}, nil}), cljs_core.Meta.X_invoke_Arity1(loc))
			}
		})
	}(&cljs_core.AFn{})
}

// Replaces the node at this loc with the value of (f node args)
// @param {...*} var_args
var Edit *cljs_core.AFn

func init() {
	Edit = func(edit *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(edit, func(loc_f_args__ ...interface{}) interface{} {
			var loc = loc_f_args__[0]
			var f = loc_f_args__[1]
			var args = cljs_core.Array_seq.X_invoke_Arity1(loc_f_args__[2:])
			_, _, _ = loc, f, args
			return Replace.X_invoke_Arity2(loc, cljs_core.Apply.X_invoke_Arity3(f, Node.X_invoke_Arity1(loc), args))
		})
	}(&cljs_core.AFn{})
}

// Inserts the item as the leftmost child of the node at this loc,
// without moving
var Insert_child *cljs_core.AFn

func init() {
	Insert_child = func(insert_child *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(insert_child, func(loc interface{}, item interface{}) interface{} {
			return Replace.X_invoke_Arity2(loc, Make_node.X_invoke_Arity3(loc, Node.X_invoke_Arity1(loc), cljs_core.Cons.X_invoke_Arity2(item, Children.X_invoke_Arity1(loc)).(*cljs_core.CljsCoreCons)))
		})
	}(&cljs_core.AFn{})
}

// Inserts the item as the rightmost child of the node at this loc,
// without moving
var Append_child *cljs_core.AFn

func init() {
	Append_child = func(append_child *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(append_child, func(loc interface{}, item interface{}) interface{} {
			return Replace.X_invoke_Arity2(loc, Make_node.X_invoke_Arity3(loc, Node.X_invoke_Arity1(loc), cljs_core.Concat.X_invoke_Arity2(Children.X_invoke_Arity1(loc), (&cljs_core.CljsCorePersistentVector{nil, 1, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{item}, nil})).(*cljs_core.CljsCoreLazySeq)))
		})
	}(&cljs_core.AFn{})
}

// Moves to the next loc in the hierarchy, depth-first. When reaching
// the end, returns a distinguished loc detectable via end?. If already
// at the end, stays there.
var Next *cljs_core.AFn

func init() {
	Next = func(next *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(next, func(loc interface{}) interface{} {
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end", Fqn: "end", X_hash: float64(-268185958)}), loc.(cljs_core.CljsCoreIFn).X_invoke_Arity1(float64(1))) {
				return loc
			} else {
				{
					var or__79537__auto__ = func() interface{} {
						var and__79525__auto__ = Branch_QMARK_.X_invoke_Arity1(loc)
						_ = and__79525__auto__
						if cljs_core.Truth_(and__79525__auto__) {
							return Down.X_invoke_Arity1(loc)
						} else {
							return and__79525__auto__
						}
					}()
					_ = or__79537__auto__
					if cljs_core.Truth_(or__79537__auto__) {
						return or__79537__auto__
					} else {
						{
							var or__79537__auto_____1 = Right.X_invoke_Arity1(loc)
							_ = or__79537__auto_____1
							if cljs_core.Truth_(or__79537__auto_____1) {
								return or__79537__auto_____1
							} else {
								{
									var p = loc
									_ = p
									for {
										if cljs_core.Truth_(Up.X_invoke_Arity1(p)) {
											{
												var or__79537__auto_____2 = Right.X_invoke_Arity1(Up.X_invoke_Arity1(p))
												_ = or__79537__auto_____2
												if cljs_core.Truth_(or__79537__auto_____2) {
													return or__79537__auto_____2
												} else {
													p = Up.X_invoke_Arity1(p)
													continue
												}
											}
										} else {
											return (&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{Node.X_invoke_Arity1(p), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end", Fqn: "end", X_hash: float64(-268185958)})}, nil})
										}
									}
								}
							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})
}

// Moves to the previous loc in the hierarchy, depth-first. If already
// at the root, returns nil.
var Prev *cljs_core.AFn

func init() {
	Prev = func(prev *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(prev, func(loc interface{}) interface{} {
			{
				var temp__4220__auto__ = Left.X_invoke_Arity1(loc)
				_ = temp__4220__auto__
				if cljs_core.Truth_(temp__4220__auto__) {
					{
						var lloc = temp__4220__auto__
						_ = lloc
						{
							var loc___1 = lloc
							_ = loc___1
							for {
								{
									var temp__4220__auto_____1 = func() interface{} {
										var and__79525__auto__ = Branch_QMARK_.X_invoke_Arity1(loc___1)
										_ = and__79525__auto__
										if cljs_core.Truth_(and__79525__auto__) {
											return Down.X_invoke_Arity1(loc___1)
										} else {
											return and__79525__auto__
										}
									}()
									_ = temp__4220__auto_____1
									if cljs_core.Truth_(temp__4220__auto_____1) {
										{
											var child = temp__4220__auto_____1
											_ = child
											loc___1 = Rightmost.X_invoke_Arity1(child)
											continue
										}
									} else {
										return loc___1
									}
								}
							}
						}
					}
				} else {
					return Up.X_invoke_Arity1(loc)
				}
			}
		})
	}(&cljs_core.AFn{})
}

// Returns true if loc represents the end of a depth-first walk
var End_QMARK_ *cljs_core.AFn

func init() {
	End_QMARK_ = func(end_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(end_QMARK_, func(loc interface{}) interface{} {
			return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end", Fqn: "end", X_hash: float64(-268185958)}), loc.(cljs_core.CljsCoreIFn).X_invoke_Arity1(float64(1)))
		})
	}(&cljs_core.AFn{})
}

// Removes the node at loc, returning the loc that would have preceded
// it in a depth-first walk.
var Remove *cljs_core.AFn

func init() {
	Remove = func(remove *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(remove, func(loc interface{}) interface{} {
			{
				var vec__42 = loc
				var node = cljs_core.Nth.X_invoke_Arity3(vec__42, float64(0), nil)
				var map__43 = cljs_core.Nth.X_invoke_Arity3(vec__42, float64(1), nil)
				var map__43___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__43) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__43)
					} else {
						return map__43
					}
				}()
				var path = map__43___1
				var l = cljs_core.Get.X_invoke_Arity2(map__43___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}))
				var ppath = cljs_core.Get.X_invoke_Arity2(map__43___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ppath", Fqn: "ppath", X_hash: float64(-1758182784)}))
				var pnodes = cljs_core.Get.X_invoke_Arity2(map__43___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "pnodes", Fqn: "pnodes", X_hash: float64(1739080565)}))
				var rs = cljs_core.Get.X_invoke_Arity2(map__43___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "r", Fqn: "r", X_hash: float64(-471384190)}))
				_, _, _, _, _, _, _, _, _ = vec__42, node, map__43, map__43___1, path, l, ppath, pnodes, rs
				if cljs_core.Nil_(path) {
					panic("Remove at top")
				} else {
					if cljs_core.Count.X_invoke_Arity1(l).(float64) > float64(0) {
						{
							var loc___1 = cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.Peek.X_invoke_Arity1(l), cljs_core.Assoc.X_invoke_ArityVariadic(path, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "l", Fqn: "l", X_hash: float64(1395893423)}), cljs_core.Pop.X_invoke_Arity1(l), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "changed?", Fqn: "changed?", X_hash: float64(-437828330)}), true)}, nil}), cljs_core.Meta.X_invoke_Arity1(loc))
							_ = loc___1
							for {
								{
									var temp__4220__auto__ = func() interface{} {
										var and__79525__auto__ = Branch_QMARK_.X_invoke_Arity1(loc___1)
										_ = and__79525__auto__
										if cljs_core.Truth_(and__79525__auto__) {
											return Down.X_invoke_Arity1(loc___1)
										} else {
											return and__79525__auto__
										}
									}()
									_ = temp__4220__auto__
									if cljs_core.Truth_(temp__4220__auto__) {
										{
											var child = temp__4220__auto__
											_ = child
											loc___1 = Rightmost.X_invoke_Arity1(child)
											continue
										}
									} else {
										return loc___1
									}
								}
							}
						}
					} else {
						return cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, 2, 5, cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{Make_node.X_invoke_Arity3(loc, cljs_core.Peek.X_invoke_Arity1(pnodes), rs), func() interface{} {
							var and__79525__auto__ = ppath
							_ = and__79525__auto__
							if cljs_core.Truth_(and__79525__auto__) {
								return cljs_core.Assoc.X_invoke_Arity3(ppath, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "changed?", Fqn: "changed?", X_hash: float64(-437828330)}), true)
							} else {
								return and__79525__auto__
							}
						}()}, nil}), cljs_core.Meta.X_invoke_Arity1(loc))
					}
				}
			}
		})
	}(&cljs_core.AFn{})
}
