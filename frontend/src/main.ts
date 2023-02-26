import './style.css';
import './app.css';
 
import {Countdown} from '../wailsjs/go/main/App';
import {EventsOn} from '../wailsjs/runtime';

// Listen for the "countdown" event and update the UI
EventsOn("countdown", (remainingTime: string) => {
    countdownResultElement.textContent = remainingTime;
});

function startCountdown() {
  const countdown = countdownInputElement.value;
  console.log(countdown);
    // Check if the input is a number
    if (isNaN(Number(countdown))) {
        alert("Please enter a number");
        return;
    }

    // Start the countdown
  Countdown(countdown);
}

// Define the HTML content
const html = `
  <div class="container">
    <h1 class="heading">Countdown Timer</h1>
    <div class="countdown-container">
      <div class="countdown-label">Enter seconds:</div>
      <div class="input-box">
        <input class="input" id="countdown-input" type="text" autocomplete="off" />
        <button class="btn" id="start-btn">Start Countdown</button>
      </div>
      <div class="countdown-label">Remaining time:</div>
      <div class="countdown" id="countdown-output">0</div>
    </div>
  </div>
`;

// Set the HTML content
document.querySelector("#app")!.innerHTML = html;

// Get the necessary DOM elements
const countdownInputElement = document.getElementById("countdown-input") as HTMLInputElement;
const countdownResultElement = document.getElementById("countdown-output")!;
const startButtonElement = document.getElementById("start-btn")!;

// Add an event listener to the start button
startButtonElement.addEventListener("click", startCountdown);

