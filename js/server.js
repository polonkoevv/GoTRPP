const express = require('express')
// const mongoose = require('mongoose')
// const path = require('path')
const exphbs = require('express-handlebars')
const authRouter = require('./authRouter')
const todoRoutes = require('./routes/Routes.js')
const bp = require('body-parser')


const PORT = 3000

const app = express()
const hbs = exphbs.create({
    defaultLayout: 'main',
    extname: 'hbs'
})


// app.use(express.urlencoded({ extended: true }))

app.engine('hbs', hbs.engine)
app.set('view engine', 'hbs')
app.set('views', './public/views')

app.use(express.static('public'))

app.use(express.json())
app.use(bp.json())
app.use(bp.urlencoded({ extended: true }))

app.use("/auth", authRouter)


// app.use(express.urlencoded({ extended: true }))


app.use(todoRoutes)

async function start() {
    try {
        // await mongoose.connect(`mongodb+srv://user:user123@actors.49unziw.mongodb.net/?retryWrites=true&w=majority`)
        app.listen(PORT)
        console.log('Server has been started...')
    } catch (e) {
      console.log(e)
    }
  }

start()