@extends('layout')

@section('content')
<style type="text/css">
   .box{
    width:600px;
    margin:0 auto;
    border:1px solid #ccc;
   }
  </style>
<div class="container box">
   <h3 align="center">Sign Up Online-Judge System</h3><br />

    @if (count($errors) > 0)
        <div class="alert alert-danger">
            <ul>
            @foreach($errors->all() as $error)
                <li>{{ $error }}</li>
            @endforeach
            </ul>
        </div>
    @endif

    <form method="post" action="{{ url('/register/store') }}">
        {{ csrf_field() }}
        <div class="form-group">
            <label>Enter Name</label>
            <input type="text" name="name" class="form-control" />
        </div>
        <div class="form-group">
            <label>Enter Email</label>
            <input type="email" name="email" class="form-control" />
        </div>
        <div class="form-group">
            <label>Enter Password</label>
            <input type="password" name="password" class="form-control" />
        </div>
        <div class="form-group">
            <label>Enter password_confirmation</label>
            <input type="password" name="password_confirmation" class="form-control" />
        </div>
        <div class="form-group">
            <input type="submit" class="btn btn-primary" value="Sign Up" />
        </div>
   </form>
</div>
@stop