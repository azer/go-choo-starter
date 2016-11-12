import choo from 'choo'
import example from './components/example'

const app = choo()
app.model(example)

export default app
export const routes = router({
  '/': example.view
})

function router (routes) {
  app.router(route => {
    const result = []

    var key
    for (key in routes) {
      result.push(route(key, routes[key]))
    }

    return result
  })

  return Object.keys(routes)
}
