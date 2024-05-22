import {
  Component,
  effect,
  ElementRef, input,
  Input, InputSignal,
  signal,
  ViewChild,
  viewChild
} from '@angular/core'
// @ts-ignore
import TypeWriter from 'typewriter-effect/dist/core'

@Component({
  selector: 'app-typewriter',
  standalone: true,
  imports: [],
  templateUrl: './typewriter.component.html',
  styleUrl: './typewriter.component.css'
})
export class TypewriterComponent {
  text: InputSignal<string | null | undefined> = input<string | null | undefined>('')
  @ViewChild('content') content: ElementRef | undefined = undefined
  writer: any = undefined

  constructor() {

    effect(() => {
      const text = this.text()
      this.type(text)
    })
  }

  type(text: string | null | undefined) {
    try {
      const el = this.content?.nativeElement
      if (el) {
        if (this.writer) {
          this.writer.stop()
        }
        this.writer = new TypeWriter(el, {
          delay: 0
        })
        this.writer.typeString(text).start();
      }
    } catch (e) {
      console.error(`typewriter error: ${e}`)
    }
  }
}
