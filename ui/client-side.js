import app from './app'

const newTree = app.start()
const oldTree = document.querySelector('main')

if (oldTree) {
  document.body.replaceChild(newTree, oldTree)
} else {
  document.body.appendChild(newTree)
}
