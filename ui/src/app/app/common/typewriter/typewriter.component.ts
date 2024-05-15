import { Component, computed, Input, signal } from '@angular/core'

@Component({
  selector: 'app-typewriter',
  standalone: true,
  imports: [],
  templateUrl: './typewriter.component.html',
  styleUrl: './typewriter.component.css'
})
export class TypewriterComponent {
  @Input() words = ['|'];
  @Input() typeSpeed = 80;
  @Input() deleteSpeed = 50;
  @Input() panelClass = '';
  complete: boolean = false;

  wordIndex = signal(0);
  word = computed(() => this.words[this.wordIndex()]);

  textIndex = signal(0);
  text = computed(() => this.word().slice(0, this.textIndex()));

  ngOnInit(): void {
      this.type();
  }

  type() {
    const int = setInterval(() => {
      this.textIndex.update((v) => v + 1);
      if (this.textIndex() > this.word().length) {
        clearInterval(int);
        if (this.words.length == 1 && this.words[0] == '|') {
          this.delete()
        } else {
          this.complete = true;
        }
      }
    }, this.typeSpeed);
  }

  delete() {
    const int = setInterval(() => {
      this.textIndex.update((v) => v - 1);
      if (this.textIndex() < 0) {
        this.wordIndex.update((v) => (v + 1) % this.words.length);
        clearInterval(int);
        this.textIndex.set(0);
        if (this.words.length == 1 && this.words[0] == '|') {
          this.type();
        }
      }
    }, this.deleteSpeed);
  }
}
