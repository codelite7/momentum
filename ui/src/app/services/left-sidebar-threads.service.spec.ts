import { TestBed } from '@angular/core/testing';

import { LeftSidebarThreadsService } from './left-sidebar-threads.service';

describe('LeftSidebarThreadsService', () => {
  let service: LeftSidebarThreadsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LeftSidebarThreadsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
