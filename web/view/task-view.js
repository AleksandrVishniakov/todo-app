import {createElement} from '../render.js';

const createTaskTemplate = (task) => {
  const {color, description, dueDate} = task;

  return (
  `<article class="card card--${color}">
    <div class="card__form">
      <div class="card__inner">
        <div class="card__textarea-wrap">
          <p class="card__text">${description}</p>
        </div>
        <div class="card__settings">
          <div class="card__date-deadline">
            <p class="card__input-deadline-wrap">
              <span class="card__date">${dueDate}</span>
            </p>
          </div>
        </div>
        <div class="card__control">
        <button type="button" class="card__btn">
          Редактировать
        </button>
      </div>
      </div>
    </div>
  </article>`
);
}

export default class TaskView {
  constructor(task) {
    this.task = task;
  }

  getTemplate() {
    return createTaskTemplate(this.task);
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
}
