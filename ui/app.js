import choo from 'choo'
import example from './components/example'
import routes from "./routes"

const app = choo()
app.model(example)
router(routes)

export default app

// We need this extra function until Choo lets us get list of the routes easily
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
