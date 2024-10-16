export class Greeting {
  constructor(target) {
    this.target = target
  }

  sayHello() {
    this.target.textContent = "Hello from JS";
  }
}
