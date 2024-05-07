import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ThreadButtonComponent } from './thread-button.component';

describe('ThreadButtonComponent', () => {
  let component: ThreadButtonComponent;
  let fixture: ComponentFixture<ThreadButtonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ThreadButtonComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ThreadButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
