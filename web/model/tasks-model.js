import {generateTask} from './task.js';

export default class TasksModel {
  tasks = Array.from({length: 4}, generateTask);

  getTasks = () => this.tasks;

  addTask(task) {
    this.tasks.push(task)
  }
}