import { NextSeo } from 'next-seo';
import Link from 'next/link';
import { useForm } from 'react-hook-form';
import IdentityLayout from '../layouts/Identity';

const SignIn = () => {
  const { handleSubmit, register, errors } = useForm();

  const onSubmit = (data) => {
    console.log('form submit >>>', data);
  };

  return (
    <>
      <NextSeo title="Wetalk | Sign In" description="Talk together!" />
      <IdentityLayout>
        <div className="w-1/4 mx-auto flex flex-col p-10 bg-white rounded-lg shadow-xl">
          <div className="mb-10 text-center">
            <h1 className="text-3xl font-bold">Welcome back!</h1>
            <p className="text-gray-500">We're so excited to see you again!</p>
          </div>
          <div className="flex flex-col">
            <form onSubmit={handleSubmit(onSubmit)}>
              <div className="mb-6">
                <label
                  className="block text-gray-700 text-sm font-bold mb-2"
                  htmlFor="username"
                >
                  Email
                </label>
                <input
                  className="appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500"
                  id="username"
                  name="username"
                  type="text"
                  placeholder="Username"
                  ref={register({
                    required: true,
                    pattern: /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
                  })}
                />
                {errors.username && (
                  <p className="text-red-500 text-xs italic mt-1">
                    Email is invalid.
                  </p>
                )}
              </div>

              <div className="mb-6">
                <label
                  className="block text-gray-700 text-sm font-bold mb-2"
                  htmlFor="password"
                >
                  Password
                </label>
                <input
                  className="appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500"
                  id="password"
                  name="password"
                  type="password"
                  placeholder="Password"
                  ref={register({ required: true })}
                />
                {errors.password && (
                  <p className="text-red-500 text-xs italic mt-1">
                    Please input password.
                  </p>
                )}
              </div>

              <div className="w-full mb-4">
                <button
                  className="w-full shadow bg-blue-500 hover:bg-blue-400 focus:outline-none text-white font-bold py-2 px-4 rounded"
                  type="submit"
                >
                  Sign In
                </button>
              </div>

              <div className="text-gray-500">
                Need an account?
                <Link href="/signup">
                  <a className="text-blue-500 ml-2 font-bold">Register</a>
                </Link>
              </div>
            </form>
          </div>
        </div>
      </IdentityLayout>
    </>
  );
};

export default SignIn;
