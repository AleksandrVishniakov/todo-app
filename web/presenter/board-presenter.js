import BoardView from '../view/board-view.js';
import TaskListView from '../view/task-list-view.js';
import TaskView from '../view/task-view.js';
import TaskEditView from '../view/task-edit-view.js';
import LoadMoreButtonView from '../view/load-more-button-view.js';
import TaskModel from '../model/task.js';
import {render} from '../render.js';

export default class BoardPresenter {
  taskListComponent = new TaskListView()
  boardComponent = new BoardView();

  addNewTask() {
    const newTask = new TaskModel()
    render(new TaskEditView(newTask), this.taskListComponent.getElement());
  }

  init = (boardContainer, tasksModel) => {
    this.boardContainer = boardContainer;
    this.tasksModel = tasksModel;
    this.boardTasks = this.tasksModel.getTasks();

    render(this.boardComponent, this.boardContainer);
    render(this.taskListComponent, this.boardComponent.getElement());

    for (let i = this.boardTasks.length - 1; i >= 0; i--) {
      const newTaskView = this.boardTasks[i].ended ? new TaskView(this.boardTasks[i]) : new TaskEditView(this.boardTasks[i])

      render(newTaskView, this.taskListComponent.getElement());

      console.log(this.boardTasks[i])
    }

    render(new LoadMoreButtonView(), this.boardComponent.getElement());
  };
}
