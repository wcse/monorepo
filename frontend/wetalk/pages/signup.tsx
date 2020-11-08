import { NextSeo } from 'next-seo';
import Link from 'next/link';
import { useForm } from 'react-hook-form';

import IdentityLayout from 'layouts/Identity';

const SignUp = () => {
  const { handleSubmit, register, errors, watch } = useForm();

  const onSubmit = (data) => {
    console.log('form submit >>>', data);
  };

  return (
    <>
      <NextSeo title="Wetalk | Sign Up" description="Talk together!" />

      <IdentityLayout>
        <div className="w-1/2 mx-auto flex flex-col p-10 bg-white rounded-lg shadow-xl">
          <h1 className="text-3xl font-bold mb-5 text-center">REGISTER</h1>
          <div className="flex">
            <div className="w-1/2 mr-5 pr-5 sign-up-form">
              <div>
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
                      id="email"
                      name="email"
                      type="text"
                      placeholder="Email"
                      ref={register({
                        required: true,
                        pattern: /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
                      })}
                    />
                    {errors.email && (
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

                  <div className="mb-6">
                    <label
                      className="block text-gray-700 text-sm font-bold mb-2"
                      htmlFor="password"
                    >
                      Confirm Password
                    </label>
                    <input
                      className="appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500"
                      id="confirm_password"
                      name="confirm_password"
                      type="password"
                      placeholder="Confirm Password"
                      ref={register({
                        validate: (value) => value === watch('password'),
                      })}
                    />
                    {errors.confirm_password && (
                      <p className="text-red-500 text-xs italic mt-1">
                        The password confirmation dose not match.
                      </p>
                    )}
                  </div>

                  <div className="w-full mb-4">
                    <button
                      className="w-full shadow bg-blue-500 hover:bg-blue-400 focus:outline-none text-white font-bold py-2 px-4 rounded"
                      type="submit"
                    >
                      Sign Up
                    </button>
                  </div>
                </form>
              </div>
              <div className="text-gray-500">
                <Link href="/signin">
                  <a className="text-blue-500 ml-2 font-bold">
                    Already have an account?
                  </a>
                </Link>
              </div>
            </div>

            <div className="w-1/2">
              <div className="flex flex-col justify-center h-full">
                <div className="w-full mb-4">
                  <button
                    className="w-full shadow bg-purple-500 hover:bg-purple-400 focus:outline-none text-white font-bold py-2 px-4 rounded"
                    type="submit"
                  >
                    Google
                  </button>
                </div>

                <div className="w-full">
                  <button
                    className="w-full shadow bg-blue-500 hover:bg-blue-400 focus:outline-none text-white font-bold py-2 px-4 rounded"
                    type="submit"
                  >
                    Facebook
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </IdentityLayout>
    </>
  );
};

export default SignUp;
