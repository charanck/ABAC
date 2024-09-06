import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

const BASE_URL = 'http://localhost:3000';
@Injectable({
  providedIn: 'root',
})
export class ResourceService {
  constructor(private http: HttpClient) {}

  listResources(pageNumber: number, pageSize: number): Observable<Resource[]> {
    return this.http
      .get<Resource[]>(`${BASE_URL}/resources`, {
        params: {
          pageNumber: pageNumber,
          pageSize: pageSize,
        },
      })
      .pipe(
        map((resources) => {
          for (let i = 0; i < resources.length; i++) {
            resources[i].deleted = new Date(resources[i].deleted);
            resources[i].created = new Date(resources[i].created);
            resources[i].updated = new Date(resources[i].updated);
          }
          return resources;
        })
      );
  }
}

export interface Resource {
  description: string;
  id: string;
  name: string;
  ownerId: string;
  policyId: string;
  updated: Date;
  created: Date;
  deleted: Date;
}
