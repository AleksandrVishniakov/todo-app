const createElement = (template) => {
  const newElement = document.createElement('div'); // 1
  newElement.innerHTML = template; // 2

  return newElement.firstElementChild; // 3
};

const render = (component, container) => {
  const element = component.getElement();
  container.append(element);
};

export {createElement, render};