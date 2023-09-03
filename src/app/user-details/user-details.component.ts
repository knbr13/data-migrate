import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.css'],
})
export class UserDetailsComponent implements OnInit {
  userId: number = 1;
  user: any;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private http: HttpClient
  ) {}

  ngOnInit(): void {
    this.route.params.subscribe((params) => {
      this.userId = +params['id'];
      this.fetchUserDetails(this.userId);
    });
  }

  fetchUserDetails(id: number): void {
    const apiUrl = `https://reqres.in/api/users/${id}`;
    this.http.get(apiUrl).subscribe((data: any) => {
      this.user = data.data;
    });
  }

  goBack(): void {
    this.router.navigate(['/users']);
  }
}
