exports.handler = null

exports.loadPlugin = function (codePath) {
  try {
    this.handler = require(codePath)
  } catch (err) {
    console.error(`load error ${err}`)
  }
}
