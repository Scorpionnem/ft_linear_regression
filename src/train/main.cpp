/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.cpp                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mbatty <mbatty@student.42angouleme.fr>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2025/12/08 13:08:53 by mbatty            #+#    #+#             */
/*   Updated: 2025/12/08 15:51:28 by mbatty           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <iostream>
#include <string>
#include <cstdlib>
#include <fstream>
#include <vector>
#include "estimatePrice.hpp"

std::pair<double, double>	parseLine(const std::string &line)
{
	std::pair<double, double>	res;

	if (line.find(',') == line.npos)
		throw std::runtime_error("Could not find ','");

	std::string	first = line.substr(0, line.find(','));
	std::string	second = line.substr(line.find(',') + 1);
	res.first = std::stod(first);
	res.second = std::stod(second);
	return (res);
}

void	gradientStep(std::vector<std::pair<double, double>> &data, double learningRate, double &theta0, double &theta1)
{
	double	sum0 = 0.0;
	double	sum1 = 0.0;
	int		m = data.size();

	for (int i = 0; i < m; i++)
	{
		double x = data[i].first;
		double y = data[i].second;
		double prediction = estimatePrice(x, theta0, theta1);
		double error = prediction - y;

		sum0 += error;
		sum1 += error * x;
	}

	double tmpTheta0 = theta0 - learningRate * (sum0 / m);
	double tmpTheta1 = theta1 - learningRate * (sum1 / m);

	theta0 = tmpTheta0;
	theta1 = tmpTheta1;
}

double	normalizeDataset(std::vector<std::pair<double, double>> &data)
{
	double	maxMileage = 0;

	for (auto &pair : data)
		if (pair.first > maxMileage)
			maxMileage = pair.first;

	for (auto &pair : data)
		pair.first /= maxMileage;
		
	return (maxMileage);
}

int	main(int ac, char **av)
{
	(void)ac;(void)av;
	if (ac != 2)
	{
		std::cerr << "Usage: ./Train *.csv" << std::endl;
		return (1);
	}
	std::ifstream	file(av[1]);
	if (!file.is_open())
	{
		std::cerr << "Error: Could not open file" << std::endl;
		return (1);
	}

	std::vector<std::pair<double, double>>	data;
	std::string	line;
	int	lineNumber = 0;

	while (getline(file, line))
	{
		lineNumber++;

		try {
			std::pair<double, double> pair = parseLine(line);
			data.push_back(pair);
		} catch (const std::exception &e)
		{
			std::cerr << "Error at line " << lineNumber << " : " << line << " (" << e.what() << ")" << std::endl;
			continue ;
		}
	}

	double maxMileage = normalizeDataset(data);

	double theta0 = 0.0;
	double theta1 = 0.0;
	double learningRate = 0.01;
	int iterations = 100000;

	for (int i = 0; i < iterations; i++)
		gradientStep(data, learningRate, theta0, theta1);

	std::cout << "theta0: " << theta0 << " theta1: " << theta1 / maxMileage << std::endl;

	return (0);
}
