import choo from "choo"
import html from "choo/html"

const view = (state, prev, send) => html`
  <main class="example">
    <h1>Title: ${state.title}</h1>
    <input type="text" oninput=${(e) => send('setTitle', e.target.value)}>
  </main>
`

export default view
