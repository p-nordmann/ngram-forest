# prefix-forest
Baseline language model based on n-grams only. 

## The Unreasonable Effectiveness of ~~Recurrent Neural Networks~~ n-gram models

How much natural language can we predict using n-gram statistics?

This repository aims to give some idea about the answer.
It provides a super simple code to learn an autoregressive model from n-gram statistics taken from a
corpus of documents.

## What is this and how do I use it?

You can have a look at [main.go](./cmd/pforest/main.go) to see how you can generate basic text from a corpus of document.
Alternatively, [p-nordmann/lorem-ipsum](https://github.com/p-nordmann/lorem-ipsum) provides a _lorem ipsum_ generator based on prefix-forest.

The main structure defined by the `pforest` package is the `Forest` structure ([forest.go](./forest.go#L11)).
It is basically a collection of prefix-trees (or tries), one for each possible
rune, memorizing the frequency of occurrence of each n-gram up to a certain depth.

In order to sample text from the prefix forest, we proceed one rune at a time.
We find the largest suffix of the context matching a known n-gram and sample
the next rune according to the statistics learnt for this n-gram.
