# Word Search

> NeetCode 150 · Backtracking · Blind 75 — DFS per starting cell, marking the path in place with an off-alphabet byte and restoring on backtrack. Go, stdlib only, tested.

## Problem

Can a word be traced through a letter grid, moving only to 4-adjacent cells and never standing on the same cell twice within the word? `ABCB` on the classic board fails precisely because of the reuse rule.

## Approach

Try every cell as a start; from each, a DFS matches one character per step and branches four ways. The visited-set is the board itself: overwrite the current cell with `'#'` (outside the alphabet, so it can't match) and restore it after the four branches return — O(1) space for path tracking, and the test asserts the board comes back intact. The mismatch check doubles as the bounds check by short-circuit order. Fail-fast matters: the first character comparison rejects most starts immediately, which is what keeps the exponential worst case theoretical on real inputs. Word Search II (dsa-217) is this DFS with a trie replacing the single word.

**Complexity:** O(cells · 3^len) worst case (no immediate backstep), O(len) recursion.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
