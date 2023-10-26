import NewTaskButtonView from '../view/new-task-button-view.js';
import {render} from './render.js';
import BoardPresenter from '../presenter/board-presenter.js';
import TasksModel from '../model/tasks-model.js';

const siteMainElement = document.querySelector('.main');
const siteHeaderElement = document.querySelector('.control');

const tasksModel = new TasksModel();
const boardPresenter = new BoardPresenter();


boardPresenter.init(siteMainElement, tasksModel);

render(new NewTaskButtonView(boardPresenter.newTask), siteHeaderElement);