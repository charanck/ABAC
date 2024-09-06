import { Component, signal, ViewChild, WritableSignal } from '@angular/core';
import { Resource, ResourceService } from './resource.service';
import { MatTableDataSource } from '@angular/material/table';
import { MatPaginator, PageEvent } from '@angular/material/paginator';

@Component({
  selector: 'app-resources',
  templateUrl: './resources.component.html',
  styleUrls: ['./resources.component.scss'],
})
export class ResourcesComponent {
  protected length = 100;
  protected pageSize = 5;
  protected pageIndex = 0;
  protected pageSizeOptions = [5, 10, 25];
  protected resources: WritableSignal<MatTableDataSource<Resource>> = signal(
    new MatTableDataSource<Resource>()
  );
  protected displayedColumns: string[] = [
    'no',
    'name',
    'ownerId',
    'policyId',
    'description',
    'menu',
  ];


  constructor(private resourceService: ResourceService) {}

  ngOnInit() {
    this.fetchResources();
  }

  handlePageEvent(e: PageEvent) {
    this.length = e.length;
    this.pageSize = e.pageSize;
    this.pageIndex = e.pageIndex;
    this.fetchResources();
  }

  fetchResources() {
    this.resourceService
      .listResources(this.pageIndex + 1, this.pageSize)
      .subscribe((resources) => {
        this.resources.set(new MatTableDataSource<Resource>(resources));
      });
  }
}
