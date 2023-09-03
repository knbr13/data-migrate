import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css'],
})
export class UserListComponent implements OnInit {
  users: any[] = [];
  searchQuery: string = '';

  constructor(private http: HttpClient, private router: Router) {}

  ngOnInit(): void {
    this.fetchUsers(1);
  }

  fetchUsers(page: number): void {
    const apiUrl = `https://reqres.in/api/users?page=${page}`;
    this.http.get(apiUrl).subscribe((data: any) => {
      this.users = data.data; // Assuming the user data is under the 'data' property
    });
  }

  navigateToUserDetails(userId: number): void {
    this.router.navigate(['user', userId]);
  }
}
