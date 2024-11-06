import { Component, inject, signal, WritableSignal } from '@angular/core';
import { Resource, ResourceService } from './resource.service';
import { MatTableDataSource } from '@angular/material/table';
import { PageEvent } from '@angular/material/paginator';
import {Router } from '@angular/router';
import {
  MAT_DIALOG_DATA,
  MatDialog,
  MatDialogActions,
  MatDialogClose,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import { CreateEditResourceDialogComponent } from './create-edit-resource-dialog/create-edit-resource-dialog.component';

@Component({
  selector: 'app-resources',
  templateUrl: './resources.component.html',
  styleUrls: ['./resources.component.scss'],
})
export class ResourcesComponent {
  private readonly dialog = inject(MatDialog);
  private readonly router = inject(Router);
  private readonly resourceService = inject(ResourceService); 

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

  viewEdit(row: Resource){
    const dialogRef = this.dialog.open(CreateEditResourceDialogComponent, {
      data: row,
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('The dialog was closed');
    });
  }

  createResource(){
    const dialogRef = this.dialog.open(CreateEditResourceDialogComponent, {
      data: null,
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('The dialog was closed');
    });
  }
}
