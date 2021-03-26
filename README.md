# README
## Hw to generate
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
