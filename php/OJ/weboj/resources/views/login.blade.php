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
    <h3 align="center">Login Online-Judge System</h3><br />

    @if(isset(Auth::user()->email))
        <script>window.location="/";</script>
    @endif

    @if ($message = Session::get('error'))
    <div class="alert alert-danger alert-block">
        <button type="button" class="close" data-dismiss="alert">×</button>
        <strong>{{ $message }}</strong>
    </div>
    @endif

    @if (count($errors) > 0)
        <div class="alert alert-danger">
            <ul>
            @foreach($errors->all() as $error)
                <li>{{ $error }}</li>
            @endforeach
            </ul>
        </div>
    @endif
  
    <form method="post" action="{{ url('/login/store') }}">
    {{ csrf_field() }}
    <div class="form-group">
        <label>Enter Email</label>
        <input type="email" name="email" class="form-control" />
    </div>
    <div class="form-group">
        <label>Enter Password</label>
        <input type="password" name="password" class="form-control" />
    </div>
    <div class="form-group">
        <input type="submit" class="btn btn-primary" value="Login" />
    </div>
    </form>
</div>
@stop