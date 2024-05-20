import { TestBed } from '@angular/core/testing';

import { EllipsisService } from './ellipsis.service';

describe('EllipsisService', () => {
  let service: EllipsisService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(EllipsisService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
