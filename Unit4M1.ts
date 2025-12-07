/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-12-07
 * @fileoverview This program is a simple calculator performing multiple operations.
 */

// welcome message and menu
console.log("Welcome to my calculator program.");
console.log("Which operation would you like to perform today? (Select by typing the letter in front of the operation.)");
console.log("a. add");
console.log("b. subtract");
console.log("c. multiply");
console.log("d. divide");
console.log("e. absolute value");
console.log("f. round");
console.log("g. raise to an exponent");
console.log("h. square root");

// get user operation
const operation: string = prompt("What operation do you want to choose: ") ?? "";

// variables
let num1: number;
let num2: number;
let result: number;

// perform operation using if statements
if (operation === "a") {
  num1 = Number(prompt("Enter the first number: "));
  num2 = Number(prompt("Enter the second number: "));
  result = num1 + num2;
  console.log(`${num1} + ${num2} = ${result}`);
} else if (operation === "b") {
  num1 = Number(prompt("Enter the first number: "));
  num2 = Number(prompt("Enter the second number: "));
  result = num1 - num2;
  console.log(`${num1} - ${num2} = ${result}`);
} else if (operation === "c") {
  num1 = Number(prompt("Enter the first number: "));
  num2 = Number(prompt("Enter the second number: "));
  result = num1 * num2;
  console.log(`${num1} * ${num2} = ${result}`);
} else if (operation === "d") {
  num1 = Number(prompt("Enter the dividend: "));
  num2 = Number(prompt("Enter the divisor: "));
  result = num1 / num2;
  console.log(`${num1} / ${num2} = ${result}`);
} else if (operation === "e") {
  num1 = Number(prompt("Enter the number to find its absolute value: "));
  result = Math.abs(num1);
  console.log(`The absolute value of ${num1} is ${result}`);
} else if (operation === "f") {
  num1 = Number(prompt("Enter the number to round: "));
  result = Math.round(num1);
  console.log(`${num1} rounded is ${result}`);
} else if (operation === "g") {
  num1 = Number(prompt("Enter the base: "));
  num2 = Number(prompt("Enter the exponent: "));
  result = Math.pow(num1, num2);
  console.log(`${num1} raised to the power of ${num2} is ${result}`);
} else if (operation === "h") {
  num1 = Number(prompt("Enter a non-negative number to find the square root: "));
  result = Math.sqrt(num1);
  console.log(`The square root of ${num1} is ${result}`);
} else {
  console.log("Invalid operation selected.");
}

console.log("\nDone.");
