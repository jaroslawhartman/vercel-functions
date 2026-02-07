# Intro

To set the tone, here’s the Vercel logo I used in this project:

![Vercel](https://jhartman.pl/files/2026-02-07-vercel-functions-with-go/vercel-logo.svg)

# vercel-functions

More thorough article on my page [https://jhartman.pl](https://jhartman.pl/posts/hosting-and-home-page/2026-02-07-vercel-functions-with-go/).

This is my hands‑on exploration of **Vercel Serverless Functions** for deploying small Go web APIs. After a bit of reshaping the folder structure (especially when I started pulling code into packages), it turned out to be straightforward and surprisingly clean.

The working structure for me looks like this:

```bash
../vercel-functions
|-- README.md
|-- api
|   |-- date.go
|   `-- hello.go
|-- cmd
|   `-- main.go
|-- go.mod
|-- pkg
|   `-- db
|       `-- db.go
|-- public
`   `-- index.html
```

Links:

- Demo: [vercel-functions-ochre.vercel.app](https://vercel-functions-ochre.vercel.app/)
- Article: [https://jhartman.pl](https://jhartman.pl/posts/hosting-and-home-page/2026-02-07-vercel-functions-with-go/)
- GitHub repo: [vercel-functions](https://github.com/jaroslawhartman/vercel-functions)
