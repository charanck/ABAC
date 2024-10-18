import { Component, signal, WritableSignal } from '@angular/core';
import { Resource, ResourceService } from './resource.service';
import { MatTableDataSource } from '@angular/material/table';
import { PageEvent } from '@angular/material/paginator';
import {Router } from '@angular/router';

@Component({
  selector: 'app-resources',
  templateUrl: './resources.component.html',
  styleUrls: ['./resources.component.scss'],
})
export class ResourcesComponent {
  protected length = 0;
  protected pageSize = 10;
  protected pageIndex = 0;
  protected pageSizeOptions = [10, 15, 20];
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


  constructor(private resourceService: ResourceService, private router:Router) {}

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
      .subscribe((response) => {
        this.length = response.pagingMetadata.total;
        this.resources.set(new MatTableDataSource<Resource>(response.data));
      });
  }

  view(row: Resource){
    this.router.navigateByUrl(`/resources/${row.id}`);
  }
}
