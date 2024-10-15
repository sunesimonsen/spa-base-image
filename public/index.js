import { Greeting } from "./Greeting.js"

const target = document.getElementById("message");
const greeting = new Greeting(target)
greeting.sayHello()
