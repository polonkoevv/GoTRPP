const Router = require('express')
const express = require('express')
const authRouter = new Router()
const controller = require('./authController')
const authMiddleware = require('./middleware/authMiddleware')
const {check} = require("express-validator")

const urlencodedParser = express.urlencoded({extended: false});

authRouter.post('/registration', urlencodedParser, controller.registration)
authRouter.post('/login', urlencodedParser, controller.login)
authRouter.get('./login', urlencodedParser, controller.toProfile)
authRouter.get('/users', authMiddleware, controller.getUsers)
// authRouter.post('/adtofavor',urlencodedParser, controller.adtofavor)

module.exports = authRouter