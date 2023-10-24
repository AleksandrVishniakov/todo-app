import NewTaskButtonView from './view/new-task-button-view.js';
import {render} from './render.js';
import BoardPresenter from './presenter/board-presenter.js';
import TasksModel from './model/tasks-model.js';

const siteMainElement = document.querySelector('.main');
const siteHeaderElement = document.querySelector('.control');

const tasksModel = new TasksModel();
const boardPresenter = new BoardPresenter();
const taskButtonView = new NewTaskButtonView()

render(taskButtonView, siteHeaderElement);
taskButtonView.onclick(() => {
    boardPresenter.addNewTask()
})

boardPresenter.init(siteMainElement, tasksModel);
