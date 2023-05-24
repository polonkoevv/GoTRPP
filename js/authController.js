// const dbase = require('./server')
const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');
const { validationResult } = require('express-validator')
const {secret} = require("./config")

// import {dbase} from './server.js'



// console.log(typeof(dbase))


const mysql = require("mysql2");

const dbase = mysql.createConnection({
    host: "127.0.0.1",
    user: "mysql",
    database: "cwork",
    password: "mysql"
});

dbase.connect(function(err){
    if (err) {
      return console.error("Ошибка: " + err.message);
    }
    else{
      console.log("Подключение к серверу MySQL успешно установлено");
    }
})

const generateAccessToken = (id, firstname, lastname, username, roles) => {
    const payload = {
        id, firstname, lastname, username, roles
    }
    return jwt.sign(payload, secret, {expiresIn: "24h"} )
}



class authController {
    
    async registration(req, res) {
        try {
            console.log(req.body)
            const {firstname,lastname,username, password} = req.body;
            // const username = "BERS"
            // const password = "BERS"
            var reslt;
            function checkQuant(callback){
            dbase.query("SELECT id FROM account WHERE username = ?", username , 
            function(err, results, fields) {
                return callback(results)
            })};
            
            checkQuant(function(result){
                reslt = result
                if (reslt != 0) {
                    return res.status(400).json({message: "Пользователь с таким именем уже существует"})
                }
                const hashPassword = bcrypt.hashSync(password, 7);
                dbase.query("INSERT INTO account (firstname,lastname,username,password) VALUES (?,?,?,?)",[firstname,lastname,username,hashPassword])
                return res.json({message: "Пользователь успешно зарегистрирован"})
            })
            
            
        } catch (e) {
            console.log(e)
            res.status(400).json({message: 'Registration error'})
        }
    }

    async login(req, res) {
        try {
            const {username, password} = req.body

            // var user = 1;

            function checkQuant(callback){
            dbase.query("SELECT id, firstname, lastname, role, username, password FROM account WHERE username = ?",[username],function (err,results,fields){
                // console.log(results)
                return callback(results);
            })}
            
            checkQuant(function(result){
                var user = result
                console.log(user)
                if (user == 0) {
                    return res.status(400).json({message: `Пользователь ${username} не найден`})
                }

                const validPassword = bcrypt.compareSync(password,user[0].password)

                if (!validPassword) {
                    return res.status(400).json({message: `Введен неверный пароль`})
                }

                console.log(user[0].lastname)

                const token = generateAccessToken(user[0].id, user[0].firstname, user[0].lastname, user[0].username, user[0].roles)
                // res.redirect('../index.html')
                console.log("Авторизация прошла успешно")
                // res.set("Authorization", "Bearer " + token)
                // return res.json({"username":username,"password":password, "Athorization":"Bearer " + token})

                res.redirect('./profile/' + token)
            })
            // console.log(user)
            

            
            // const validPassword = bcrypt.compareSync(password, user.password)
            
        } catch (e) {
            console.log(e)
            res.status(400).json({message: 'Login error'})
        }

    }

    async toProfile(req,res){
        try {
            res.send("ok")
        } catch (e) {
            console.log(e)
        }
    }

    async getUsers(req, res) {
        try {
            res.json({message: "ok"})
        } catch (e) {
            console.log(e)
        }
    }
}

module.exports = new authController()