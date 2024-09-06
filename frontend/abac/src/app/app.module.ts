import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ResourcesComponent } from './resources/resources.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatIconModule } from '@angular/material/icon';
import { ResourceService } from './resources/resource.service';
import { HttpClientModule } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatTableModule } from '@angular/material/table';
import {MatMenuModule} from '@angular/material/menu';

@NgModule({
  declarations: [AppComponent, ResourcesComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatIconModule,
    HttpClientModule,
    CommonModule,
    MatButtonModule,
    MatTableModule,
    MatPaginatorModule,
    MatMenuModule
  ],
  providers: [ResourceService],
  bootstrap: [AppComponent],
})
export class AppModule {}
