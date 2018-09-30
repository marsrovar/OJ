@extends('layout')

@section('content')
<section class="container">
    @if(Auth::user()->class == 0)
        <a href="{{ url('questionpool/create') }}" role="btn" class="btn btn-primary">New</a>
    @endif
    <table class="table table-hover">
    <thead class="thead-light">
    <tr>
        <th scope="col">Num</th>
        <th scope="col">Title</th>
        <th scope="col">Content</th>
    @if(Auth::user()->class == 0)
        <th scope="col">Edit</th>
        <th scope="col">Delete</th>
    @else
        <th scope="col"> </th>
        <th scope="col">Status</th>
    @endif
    </tr>
    </thead>
    @forelse ($data as $var)
    <tr>
        <!-- <td>{{ $var->id }}</td> -->
        <th scope="row">{{ $var->num }}</th>
        <td><a href="{{ url('questionpool/'.$var->id) }}">{{ $var->title }}</a></td>
        <td ><a href="{{ url('questionpool/'.$var->id) }}">{{ $var->content }}</a></td>
    @if(Auth::user()->class == 0)
        <td><a href="{{ url('questionpool/'.$var->id.'/edit') }}" role="btn" class="btn btn-success btn-sm">edit</a></td>
        <td><a href="{{ url('questionpool/'.$var->id) }}" role="btn" class="btn btn-danger btn-sm">delete</a></td>
    @else
        <td><a href="{{ url('questionpool/'.$var->id) }}" role="btn" class="btn btn-danger btn-sm">delete</a></td>
    @endif
    </tr>
    @empty
        <div class="alert alert-primary" role="alert">
        No Data!
        </div>
    @endforelse
    </table>
</section>
@stop