# README
## How to generate
```
# nextjs-sample
$ yarn build && yarn start

$ ls out/index.html
```
```
# nuxtjs-sample
$ yarn generate && yarn start

$ ls dist/index.html
```
```
# vite-vue-ts
$ yarn build && yarn serve

$ ls dist/index.html
```
```
# vite-vue-react
$ yarn build && yarn serve

$ ls dist/index.html
```

You can simply serve with a HTTP server for each project.
```
$ (cd out && python3 -m http.server 8000)
...

$ (cd dist && python3 -m http.server 8000)
...
```

## How I set up
Here is how I make scaffolding each project.

### nextjs
https://nextjs.org/docs/getting-started

```
$ yarn create next-app
...
# wait for a few minutes 

What is your project named? › nextjs-sample
...

$ cd nextjs-sample
$ yarn build && yarn start
```

### nuxtjs
https://nuxtjs.org/docs/2.x/get-started/installation/

```
$ mkdir nuxtjs-sample
...
create-nuxt-app v3.6.0
✨  Generating Nuxt.js project in .
? Project name: nuxtjs-sample
? Programming language: TypeScript
? Package manager: Yarn
? UI framework: Vuetify.js
? Nuxt.js modules: Progressive Web App (PWA)
? Linting tools: ESLint, Prettier, StyleLint
? Testing framework: None
? Rendering mode: Universal (SSR / SSG)
? Deployment target: Static (Static/Jamstack hosting)
? Development tools: (Press <space> to select, <a> to toggle all, <i> to invert selection)
? Continuous integration: None
? Version control system: None
...

$ yarn generate && yarn start
```

### vite vue-ts
https://vitejs.dev/guide/#browser-support

```
$ yarn create @vitejs/app
...
✔ Project name: · vite-vue-ts

Scaffolding project in /Users/tsakuma/static-site-generation-sample/vite-react-ts...
? Select a template: …
  vanilla
  vue
❯ vue-ts
  react
  react-ts
  preact
  preact-ts
  lit-element
  lit-element-ts
  svelte
  svelte-ts

...

$ cd vite-vue-ts
$ yarn && yarn dev
```

### vite react-ts
https://vitejs.dev/guide/#browser-support

```
$ yarn create @vitejs/app
...
✔ Project name: · vite-react-ts

? Select a template: …
...
❯ react-ts
...

$ cd vite-react-ts
$ yarn && yarn dev
```
