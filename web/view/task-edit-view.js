import {createElement} from '../render.js';

const createTaskEditTemplate = (task = {}) => {
  const {
    color = 'black',
    description = '',
    dueDate = null,
  } = task;
  
  return(
  `<article class="card card--edit card--${color}">
    <form class="card__form" method="get">
      <div class="card__inner">
        <div class="card__textarea-wrap">
          <label>
            <textarea
              class="card__text"
              placeholder="Текст вашей задачи пишите тут..."
              name="text"
            >${description}</textarea>
          </label>
        </div>
        <div class="card__settings">
          <div class="card__details">
              <fieldset class="card__date-deadline">
                <label class="card__input-deadline-wrap">
                  <input
                    class="card__date"
                    type="date"
                    placeholder=""
                    name="date"
                    value=${dueDate}
                  />
                </label>
              </fieldset>
          <div class="card__colors-inner">
            <h3 class="card__colors-title">Цвет</h3>
            <div class="card__colors-wrap">
              <div class="card__color">
              <input
                type="radio"
                id="color-black-4"
                class="card__color-input card__color-input--black visually-hidden"
                name="color"
                value="black"
              />
              <label
                for="color-black-4"
                class="card__color-value card__color-value--black"
                >Черный
              </label>
              </div>
              <div class="card__color">
              <input
                type="radio"
                id="color-yellow-4"
                class="card__color-input card__color-input--yellow visually-hidden"
                name="color"
                value="yellow"
                checked
              />
              <label
                for="color-yellow-4"
                class="card__color-value card__color-value--yellow"
                >Желтый</label
              >
              </div>
              <div class="card__color">
              <input
                type="radio"
                id="color-blue-4"
                class="card__color-input card__color-input--blue visually-hidden"
                name="color"
                value="blue"
              />
              <label
                for="color-blue-4"
                class="card__color-value card__color-value--blue"
                >Синий
              </label>
              </div>
              <div class="card__color">
              <input
                type="radio"
                id="color-green-4"
                class="card__color-input card__color-input--green visually-hidden"
                name="color"
                value="green"
              />
              <label
                for="color-green-4"
                class="card__color-value card__color-value--green"
                >Зеленый
              </label>
              </div>
              <div class="card__color">
              <input
                type="radio"
                id="color-pink-4"
                class="card__color-input card__color-input--pink visually-hidden"
                name="color"
                value="pink"
              />
              <label
                for="color-pink-4"
                class="card__color-value card__color-value--pink"
                >Розовый
              </label>
              <div>
            </div>
          </div>
        </div>

        <div class="card__status-btns">
          <button class="card__save" type="submit">Сохранить</button>
          <button class="card__delete" type="button">Удалить</button>
        </div>
      </div>
    </form>
  </article>`
);
}

export default class TaskEditView {
  getTemplate() {
    return createTaskEditTemplate();
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
