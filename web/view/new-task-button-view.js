import {createElement} from '../render.js';

const createNewTaskButtonTemplate = () => '<button class="control__button">+ ДОБАВИТЬ ЗАДАЧУ</button>';

export default class NewTaskButtonView {
  getTemplate() {
    this.element = createNewTaskButtonTemplate()
    return createNewTaskButtonTemplate();
  }

  getElement() {
    if (!this.element) {
      this.element = createElement(this.getTemplate());
    }

    return this.element;
  }

  onclick(callback) {
    this.element.addEventListener('click', callback)
  }

  removeElement() {
    this.element = null;
  }
}
