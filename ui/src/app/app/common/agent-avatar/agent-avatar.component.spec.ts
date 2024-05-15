import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AgentAvatarComponent } from './agent-avatar.component';

describe('AgentAvatarComponent', () => {
  let component: AgentAvatarComponent;
  let fixture: ComponentFixture<AgentAvatarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AgentAvatarComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(AgentAvatarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
