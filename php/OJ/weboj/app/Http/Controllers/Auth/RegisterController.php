<?php

namespace App\Http\Controllers\Auth;

use App\User;
use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Validator;
use Illuminate\Foundation\Auth\RegistersUsers;
use Illuminate\Http\Request;
use Auth;

class RegisterController extends Controller
{
    /*
    |--------------------------------------------------------------------------
    | Register Controller
    |--------------------------------------------------------------------------
    |
    | This controller handles the registration of new users as well as their
    | validation and creation. By default this controller uses a trait to
    | provide this functionality without requiring any additional code.
    |
    */

    use RegistersUsers;

    // /**
    //  * Where to redirect users after registration.
    //  *
    //  * @var string
    //  */
    // protected $redirectTo = '/';

    // /**
    //  * Create a new controller instance.
    //  *
    //  * @return void
    //  */
    // public function __construct()
    // {
    //     $this->middleware('guest');
    // }

    public function index()
    {   
        return view('register');
    }

    function store(Request $request)
    {
        $this->validate($request, [
            'name' => 'required|alpha|min:2',
            'email' =>' required|email|unique:users',
            'password' => 'required|alpha_num|between:6,12|confirmed',
            'password_confirmation' => 'required|alpha_num|between:6,12'
        ]);

        $user_data = array(
            'name' => $request->get('name'),
            'email' => $request->get('email'),
            'password' => $request->get('password'),
            'password_confirmation' => $request->get('password_confirmation'),
            'class' => 1,
            'status' => 0,
        );

        $this->create($user_data);
        return redirect('/');
    
    }

    /**
     * Get a validator for an incoming registration request.
     *
     * @param  array  $data
     * @return \Illuminate\Contracts\Validation\Validator
     */
    protected function validator(array $data)
    {
        return Validator::make($data, [
            'name' => 'required|alpha|min:2',
            'email' =>' required|email|unique:users',
            'password' => 'required|alpha_num|between:6,12|confirmed',
            'password_confirmation' => 'required|alpha_num|between:6,12',
        ]);
    }

    /**
     * Create a new user instance after a valid registration.
     *
     * @param  array  $data
     * @return \App\User
     */
    protected function create(array $data)
    {
        User::create([
            'name' => $data['name'],
            'email' => $data['email'],
            'password' => Hash::make($data['password']),
            'class' => $data['class'],
            'status' => $data['status'],
        ]);
    }
}