const nav = document.querySelector(".top-nav ul");
const hamburger = document.querySelector(".hamburger");
const head = document.querySelector("header");
const k1 = document.querySelector(".kapan");
const k2 = document.querySelector(".kapan1");
const k3 = document.querySelector(".kapan2");
const k4 = document.querySelector(".kapan3");

hamburger.addEventListener("click", () => {
  nav.classList.toggle("show");
  head.classList.toggle("uzunluk");
  hamburger.classList.toggle("siyah");
});
k1.addEventListener("click", () => {
  nav.classList.remove("show");
  head.classList.remove("uzunluk");
});

k2.addEventListener("click", () => {
  nav.classList.remove("show");
  head.classList.remove("uzunluk");
});

k3.addEventListener("click", () => {
  nav.classList.remove("show");
  head.classList.remove("uzunluk");
});
k4.addEventListener("click", () => {
  nav.classList.remove("show");
  head.classList.remove("uzunluk");
});
