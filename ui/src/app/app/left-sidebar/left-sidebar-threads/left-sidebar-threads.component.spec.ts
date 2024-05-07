import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LeftSidebarThreadsComponent } from './left-sidebar-threads.component';

describe('LeftSidebarThreadsComponent', () => {
  let component: LeftSidebarThreadsComponent;
  let fixture: ComponentFixture<LeftSidebarThreadsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LeftSidebarThreadsComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(LeftSidebarThreadsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
