import { Component } from '@angular/core';
import { BaseComponent } from '../base/base.component'

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [
    BaseComponent
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {

}