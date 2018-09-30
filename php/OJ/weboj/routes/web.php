<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function () {
    return view('welcome');
})->name('login');

Route::get('/login', 'LogInOutController@index');
Route::post('/login/store', 'LogInOutController@store');
Route::get('/logout', 'LogInOutController@logout');

Route::get('/register', 'Auth\RegisterController@index');
Route::post('/register/store', 'Auth\RegisterController@store');

Route::group(['middleware'  => 'auth'], function(){
    Route::resource('/questionpool','QuestionPoolController', ['except' => ['store']]);
    Route::post('/questionpool/create', 'QuestionPoolController@store');
});