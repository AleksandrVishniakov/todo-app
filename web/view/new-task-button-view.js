import {createElement} from '../src/render.js';

const createNewTaskButtonTemplate = () => '<button class="control__button">+ ДОБАВИТЬ ЗАДАЧУ</button>';

export default class NewTaskButtonView {

  handleClick = null;

  constructor(callback) {
    this.handleClick = callback;
    this.getElement().addEventListener('click', this.clickHandler);
  }

  getTemplate() {
    return createNewTaskButtonTemplate();
  }

  getElement() {
    if (!this.element) {
      this.element = createElement(this.getTemplate());
    }

    return this.element;
  }

  removeElement() {
    this.element = null;
  }

  clickHandler = (evt) => {
    evt.preventDefault();
    this.handleClick();
  };
}
