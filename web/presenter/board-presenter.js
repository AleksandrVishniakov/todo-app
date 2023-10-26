import BoardView from '../view/board-view.js';
import TaskListView from '../view/task-list-view.js';
import TaskView from '../view/task-view.js';
import TaskEditView from '../view/task-edit-view.js';
import LoadMoreButtonView from '../view/load-more-button-view.js';
import {render, unrender} from '../src/render.js';
import { generateTask } from '../model/task.js';

export default class BoardPresenter {
  boardComponent = new BoardView();
  taskListComponent = new TaskListView();

  init = (boardContainer, tasksModel) => {
    this.boardContainer = boardContainer;
    this.tasksModel = tasksModel;
    this.boardTasks = [...this.tasksModel.getTasks()];

    render(this.boardComponent, this.boardContainer);
    render(this.taskListComponent, this.boardComponent.getElement());

    for (let i = 0; i < this.boardTasks.length; i++) {
      render(new TaskView(this.boardTasks[i], this.editHandler), this.taskListComponent.getElement());
    }

    render(new LoadMoreButtonView(), this.boardComponent.getElement());
  };

  newTask = () => {
    this.tasksModel.addTask(generateTask());
    render(new TaskEditView(this.boardTasks[this.boardTasks.length], this.submitHandler, this.deleteHandler), this.taskListComponent.getElement());
  };

  deleteHandler = (target) => {
    this.tasksModel.removeTask(this.boardTasks.indexOf(target));
    unrender(target);
  }

  submitHandler = (target) => {
    if(!target.task) target.task = {};
    target.task.color = target.getElement().querySelector('[name="color"]:checked').value;
    target.task.description = target.getElement().querySelector('.card__text').value
    target.task.dueDate = target.getElement().querySelector('.card__date').value
    console.log(target)
    unrender(target)
    render(new TaskView(target.task, this.editHandler), this.taskListComponent.getElement());
  }

  editHandler = (target) => {
    unrender(target)
    console.log(target)
    render(new TaskEditView(target.task, this.submitHandler, this.deleteHandler), this.taskListComponent.getElement());
  }
}